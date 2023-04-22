package hashgenerator

type HashGenerator interface {
	GenerateHash(string) string
}
