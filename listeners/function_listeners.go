package listeners

//func FunctionTypeListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
//	fmt.Println("inside FunctionTypeListener ... .......")
//	children := antlrCtx.GetChildren()
//	for _, v := range children {
//		switch parserContext := v.(type) {
//		case *parser.VarSpecContext:
//			{
//				handler.VarSpecContextHandler(parserContext)
//			}
//		}
//	}
//	return nil
//}
//
//func PrimaryExprListener(goListener *GoListener, antlrCtx antlr.ParserRuleContext) error {
//	//	fmt.Println("inside PrimaryExprListener... .......")
//	children := antlrCtx.GetChildren()
//	for _, v := range children {
//		fmt.Println("1010101101100 ***")
//		fmt.Printf("%T\n", v)
//		fmt.Printf("%+v\n", v)
//		//fmt.Println("1010101101100")
//		switch parserContext := v.(type) {
//		case *parser.OperandContext:
//			{
//				handler.OperandContextHandler(parserContext)
//			}
//		case *parser.ArgumentsContext:
//			{
//				handler.ArgumentsContextHandler(parserContext)
//			}
//		case *parser.PrimaryExprContext:
//			{
//				handler.PrimaryExprContextHandler(parserContext)
//			}
//		}
//	}
//	return nil
//}
