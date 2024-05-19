package processor

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func getResultA(ctx context.Context, in int) (int, error) {
	d := time.Duration(10+rand.Intn(10)) * time.Millisecond
	time.Sleep(d)
	v := int(d) + in
	return v, nil
}

func getResultB(ctx context.Context, in int) (int, error) {
	d := time.Duration(10+rand.Intn(10)) * time.Millisecond
	time.Sleep(d)
	v := int(d) + in
	return v, nil
}

func getResultC(ctx context.Context, in int) (string, error) {
	d := time.Duration(5+rand.Intn(10)) * time.Millisecond
	time.Sleep(d)
	v := fmt.Sprintf("%d", in)
	return v, nil
}

type Processor struct {
	OutA chan int
	OutB chan int
	OutC chan string
	InC  chan int
	Errs chan error
}

func (p *Processor) Launch(ctx context.Context, inA int, inB int) {
	go func() {
		aOut, err := getResultA(ctx, inA)
		if err != nil {
			p.Errs <- err
			return
		}
		p.OutA <- aOut
	}()
	go func() {
		bOut, err := getResultB(ctx, inB)
		if err != nil {
			p.Errs <- err
			return
		}
		p.OutB <- bOut
	}()
	go func() {
		select {
		case <-ctx.Done():
			return
		case inputC := <-p.InC:
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.Errs <- err
				return
			}
			p.OutC <- cOut
		}
	}()
}

func (p *Processor) WaitForAB(ctx context.Context) (int, error) {
	var inputC int
	count := 0
	for count < 2 {
		select {
		case a := <-p.OutA:
			inputC += a
			count++
		case b := <-p.OutB:
			inputC += b
			count++
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}
	return inputC, nil
}

func (p *Processor) WaitForC(ctx context.Context) (string, error) {
	select {
	case c := <-p.OutC:
		return c, nil
	case err := <-p.Errs:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
