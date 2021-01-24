package servant

type MiQaServerImp struct {
}

//EchoHello is a test function
func (imp *MiQaServerImp) EchoHello(name string, greeting *string) (int32, error) {
	*greeting = "Hello " + name
	return 0, nil
}
