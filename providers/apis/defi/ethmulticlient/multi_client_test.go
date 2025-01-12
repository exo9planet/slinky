package ethmulticlient_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/providers/apis/defi/ethmulticlient"
	"github.com/skip-mev/slinky/providers/apis/defi/ethmulticlient/mocks"
)

func TestMultiClient(t *testing.T) {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	testcases := []struct {
		name            string
		client          ethmulticlient.EVMClient
		args            []rpc.BatchElem
		expectedResults []interface{}
		err             error
	}{
		{
			name: "no elems, no-ops",
			client: ethmulticlient.NewMultiRPCClient(
				[]ethmulticlient.EVMClient{},
				[]config.Endpoint{},
				logger,
			),
			args: []rpc.BatchElem{},
			err:  nil,
		},
		{
			name: "single client failure height request",
			client: ethmulticlient.NewMultiRPCClient(
				[]ethmulticlient.EVMClient{
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"", ""},
						[]error{nil, fmt.Errorf("height req failed")},
					),
				},
				[]config.Endpoint{},
				logger,
			),
			args: []rpc.BatchElem{{}},
			err:  fmt.Errorf("endpoint request failed"),
		},
		{
			name: "single client failure hex height decode",
			client: ethmulticlient.NewMultiRPCClient(
				[]ethmulticlient.EVMClient{
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"", "zzzzzz"},
						[]error{nil, nil},
					),
				},
				[]config.Endpoint{},
				logger,
			),
			args: []rpc.BatchElem{{}},
			err:  fmt.Errorf("could not decode hex eth height"),
		},
		{
			name: "single client success",
			client: ethmulticlient.NewMultiRPCClient(
				[]ethmulticlient.EVMClient{
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"some value", "0x12c781c"},
						[]error{nil, nil},
					),
				},
				[]config.Endpoint{{URL: "foobar"}},
				logger,
			),
			args:            []rpc.BatchElem{{}},
			expectedResults: []interface{}{"some value"},
			err:             nil,
		},
		{
			name: "two clients one failed height request",
			client: ethmulticlient.NewMultiRPCClient(
				[]ethmulticlient.EVMClient{
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"", ""},
						[]error{nil, fmt.Errorf("height req failed")},
					),
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"some value", "0x12c781c"},
						[]error{nil, nil},
					),
				},
				[]config.Endpoint{{URL: "foobar"}, {URL: "baz"}},
				logger,
			),
			args:            []rpc.BatchElem{{}},
			expectedResults: []interface{}{"some value"},
			err:             nil,
		},
		{
			name: "two clients different heights",
			client: ethmulticlient.NewMultiRPCClient(
				[]ethmulticlient.EVMClient{
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"value1", "0x12c781b"},
						[]error{nil, nil},
					),
					createEVMClientWithResponse(
						t,
						nil,
						[]string{"value2", "0x12c781c"},
						[]error{nil, nil},
					),
				},
				[]config.Endpoint{{URL: "foobar"}, {URL: "baz"}},
				logger,
			),
			args:            []rpc.BatchElem{{}},
			expectedResults: []interface{}{"value2"},
			err:             nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.client.BatchCallContext(context.TODO(), tc.args)
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
				for i, result := range tc.expectedResults {
					require.Equal(t, result, *tc.args[i].Result.(*string))
				}
			}
		})
	}
}

func createEVMClientWithResponse(
	t *testing.T,
	failedRequestErr error,
	responses []string,
	errs []error,
) ethmulticlient.EVMClient {
	t.Helper()

	c := mocks.NewEVMClient(t)
	if failedRequestErr != nil {
		c.On("BatchCallContext", mock.Anything, mock.Anything).Return(failedRequestErr)
	} else {
		c.On("BatchCallContext", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			elems, ok := args.Get(1).([]rpc.BatchElem)
			require.True(t, ok)

			require.True(t, ok)
			require.Equal(t, len(elems), len(responses))
			require.Equal(t, len(elems), len(errs))

			for i, elem := range elems {
				elem.Result = &responses[i]
				elem.Error = errs[i]
				elems[i] = elem
			}
		})
	}

	return c
}
