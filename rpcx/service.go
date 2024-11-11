package service

import "context"

type Args struct {
	A, B int
}
type Reply struct {
	C int
}

type Arith int

func (t *Arith) Multiply(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}
