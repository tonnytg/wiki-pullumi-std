package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudfunctions"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Nome da Cloud Function
		functionName := "my-cloud-function"

		// Configuração da Cloud Function
		function, err := cloudfunctions.NewFunction(ctx, functionName, &cloudfunctions.FunctionArgs{
			Runtime:       "go111",                                         // Pode ser alterado para a versão desejada do Go.
			SourceArchive: pulumi.String("path/to/your/function-code.zip"), // Caminho para o código da função.
			EntryPoint:    pulumi.String("YourFunctionEntryPoint"),         // Ponto de entrada para a função.
			TriggerHttp:   pulumi.Bool(true),                               // Configuração de trigger HTTP.
		})
		if err != nil {
			return err
		}

		// Expondo a URL da Cloud Function após a criação
		ctx.Export("functionUrl", function.HttpsTriggerUrl)

		return nil
	})
}
