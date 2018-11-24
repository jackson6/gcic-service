package grpc

import (
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// rpcError defines the status from an RPC.
type rpcError struct {
	code codes.Code
	desc string
}

func (e *rpcError) Error() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s", e.code, e.desc)
}

// toRPCErr converts an error into a rpcError.
func toRPCErr(err error) error {
	switch err.(type) {
	case *rpcError:
		return err
	default:
		switch err {
		case context.DeadlineExceeded:
			return &rpcError{
				code: codes.DeadlineExceeded,
				desc: err.Error(),
			}
		case context.Canceled:
			return &rpcError{
				code: codes.Canceled,
				desc: err.Error(),
			}
		case grpc.ErrClientConnClosing:
			return &rpcError{
				code: codes.FailedPrecondition,
				desc: err.Error(),
			}
		}

	}
	return grpc.Errorf(codes.Unknown, "%v", err)
}

// convertCode converts a standard Go error into its canonical code. Note that
// this is only used to translate the error returned by the server applications.
func convertCode(err error) codes.Code {
	switch err {
	case nil:
		return codes.OK
	case io.EOF:
		return codes.OutOfRange
	case io.ErrClosedPipe, io.ErrNoProgress, io.ErrShortBuffer, io.ErrShortWrite, io.ErrUnexpectedEOF:
		return codes.FailedPrecondition
	case os.ErrInvalid:
		return codes.InvalidArgument
	case context.Canceled:
		return codes.Canceled
	case context.DeadlineExceeded:
		return codes.DeadlineExceeded
	}
	switch {
	case os.IsExist(err):
		return codes.AlreadyExists
	case os.IsNotExist(err):
		return codes.NotFound
	case os.IsPermission(err):
		return codes.PermissionDenied
	}
	return codes.Unknown
}

func wait(ctx context.Context) bool {
	if ctx == nil {
		return false
	}
	wait, ok := ctx.Value("wait").(bool)
	if !ok {
		return false
	}
	return wait
}
