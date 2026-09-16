# 颜色与尺寸 Token：完整声明

[用途与示例](/zh-cn/api/theme) · [全部 API 索引](/zh-cn/api/)

以下声明由 Go AST 提取，保留源码英文契约和字段注释，隐藏私有字段；`struct {}` 表示外部不可直接配置的内部表示，不表示可以用零值替代构造函数。常量块保留 iota 上下文。

## colorTokenNamespace {#api-colortokennamespace}

```text
colorTokenNamespace provides discoverable, typed color tokens without
exposing stringly-typed names to applications.
```


```go
type colorTokenNamespace struct {
	Primitive primitiveColorTokens
	Semantic  semanticColorTokens
}
```

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:5`

## primitiveColorTokens {#api-primitivecolortokens}


```go
type primitiveColorTokens struct {
	White, Black                                                                                                                      ColorToken
	Red50, Red100, Red200, Red300, Red400, Red500, Red600, Red700, Red800, Red900, Red950                                             ColorToken
	Orange50, Orange100, Orange200, Orange300, Orange400, Orange500, Orange600, Orange700, Orange800, Orange900, Orange950            ColorToken
	Amber50, Amber100, Amber200, Amber300, Amber400, Amber500, Amber600, Amber700, Amber800, Amber900, Amber950                       ColorToken
	Yellow50, Yellow100, Yellow200, Yellow300, Yellow400, Yellow500, Yellow600, Yellow700, Yellow800, Yellow900, Yellow950            ColorToken
	Lime50, Lime100, Lime200, Lime300, Lime400, Lime500, Lime600, Lime700, Lime800, Lime900, Lime950                                  ColorToken
	Green50, Green100, Green200, Green300, Green400, Green500, Green600, Green700, Green800, Green900, Green950                       ColorToken
	Emerald50, Emerald100, Emerald200, Emerald300, Emerald400, Emerald500, Emerald600, Emerald700, Emerald800, Emerald900, Emerald950 ColorToken
	Teal50, Teal100, Teal200, Teal300, Teal400, Teal500, Teal600, Teal700, Teal800, Teal900, Teal950                                  ColorToken
	Cyan50, Cyan100, Cyan200, Cyan300, Cyan400, Cyan500, Cyan600, Cyan700, Cyan800, Cyan900, Cyan950                                  ColorToken
	Sky50, Sky100, Sky200, Sky300, Sky400, Sky500, Sky600, Sky700, Sky800, Sky900, Sky950                                             ColorToken
	Blue50, Blue100, Blue200, Blue300, Blue400, Blue500, Blue600, Blue700, Blue800, Blue900, Blue950                                  ColorToken
	Indigo50, Indigo100, Indigo200, Indigo300, Indigo400, Indigo500, Indigo600, Indigo700, Indigo800, Indigo900, Indigo950            ColorToken
	Violet50, Violet100, Violet200, Violet300, Violet400, Violet500, Violet600, Violet700, Violet800, Violet900, Violet950            ColorToken
	Purple50, Purple100, Purple200, Purple300, Purple400, Purple500, Purple600, Purple700, Purple800, Purple900, Purple950            ColorToken
	Fuchsia50, Fuchsia100, Fuchsia200, Fuchsia300, Fuchsia400, Fuchsia500, Fuchsia600, Fuchsia700, Fuchsia800, Fuchsia900, Fuchsia950 ColorToken
	Pink50, Pink100, Pink200, Pink300, Pink400, Pink500, Pink600, Pink700, Pink800, Pink900, Pink950                                  ColorToken
	Rose50, Rose100, Rose200, Rose300, Rose400, Rose500, Rose600, Rose700, Rose800, Rose900, Rose950                                  ColorToken
	Slate50, Slate100, Slate200, Slate300, Slate400, Slate500, Slate600, Slate700, Slate800, Slate900, Slate950                       ColorToken
	Gray50, Gray100, Gray200, Gray300, Gray400, Gray500, Gray600, Gray700, Gray800, Gray900, Gray950                                  ColorToken
	Zinc50, Zinc100, Zinc200, Zinc300, Zinc400, Zinc500, Zinc600, Zinc700, Zinc800, Zinc900, Zinc950                                  ColorToken
	Neutral50, Neutral100, Neutral200, Neutral300, Neutral400, Neutral500, Neutral600, Neutral700, Neutral800, Neutral900, Neutral950 ColorToken
	Stone50, Stone100, Stone200, Stone300, Stone400, Stone500, Stone600, Stone700, Stone800, Stone900, Stone950                       ColorToken
	Taupe50, Taupe100, Taupe200, Taupe300, Taupe400, Taupe500, Taupe600, Taupe700, Taupe800, Taupe900, Taupe950                       ColorToken
	Mauve50, Mauve100, Mauve200, Mauve300, Mauve400, Mauve500, Mauve600, Mauve700, Mauve800, Mauve900, Mauve950                       ColorToken
	Mist50, Mist100, Mist200, Mist300, Mist400, Mist500, Mist600, Mist700, Mist800, Mist900, Mist950                                  ColorToken
	Olive50, Olive100, Olive200, Olive300, Olive400, Olive500, Olive600, Olive700, Olive800, Olive900, Olive950                       ColorToken
}
```

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:10`

## semanticColorTokens {#api-semanticcolortokens}


```go
type semanticColorTokens struct {
	Surface, SurfaceHigh, Text, Accent, AccentHover, OnAccent, FocusRing, Danger, Success, Info, Warn, Border, Shadow ColorToken
}
```

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:40`

## Color {#api-color}

```text
Color is the public color-token namespace.
```


```go
var Color colorTokenNamespace
```

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:45`

## ColorPrimitivePaletteWhite {#api-colorprimitivepalettewhite}


```go
const (
	ColorPrimitivePaletteWhite ColorToken = "primitive.palette.white"
	ColorPrimitivePaletteBlack ColorToken = "primitive.palette.black"
	ColorPrimitiveRed50        ColorToken = "primitive.red.50"
	ColorPrimitiveRed100       ColorToken = "primitive.red.100"
	ColorPrimitiveRed200       ColorToken = "primitive.red.200"
	ColorPrimitiveRed300       ColorToken = "primitive.red.300"
	ColorPrimitiveRed400       ColorToken = "primitive.red.400"
	ColorPrimitiveRed500       ColorToken = "primitive.red.500"
	ColorPrimitiveRed600       ColorToken = "primitive.red.600"
	ColorPrimitiveRed700       ColorToken = "primitive.red.700"
	ColorPrimitiveRed800       ColorToken = "primitive.red.800"
	ColorPrimitiveRed900       ColorToken = "primitive.red.900"
	ColorPrimitiveRed950       ColorToken = "primitive.red.950"
	ColorPrimitiveOrange50     ColorToken = "primitive.orange.50"
	ColorPrimitiveOrange100    ColorToken = "primitive.orange.100"
	ColorPrimitiveOrange200    ColorToken = "primitive.orange.200"
	ColorPrimitiveOrange300    ColorToken = "primitive.orange.300"
	ColorPrimitiveOrange400    ColorToken = "primitive.orange.400"
	ColorPrimitiveOrange500    ColorToken = "primitive.orange.500"
	ColorPrimitiveOrange600    ColorToken = "primitive.orange.600"
	ColorPrimitiveOrange700    ColorToken = "primitive.orange.700"
	ColorPrimitiveOrange800    ColorToken = "primitive.orange.800"
	ColorPrimitiveOrange900    ColorToken = "primitive.orange.900"
	ColorPrimitiveOrange950    ColorToken = "primitive.orange.950"
	ColorPrimitiveAmber50      ColorToken = "primitive.amber.50"
	ColorPrimitiveAmber100     ColorToken = "primitive.amber.100"
	ColorPrimitiveAmber200     ColorToken = "primitive.amber.200"
	ColorPrimitiveAmber300     ColorToken = "primitive.amber.300"
	ColorPrimitiveAmber400     ColorToken = "primitive.amber.400"
	ColorPrimitiveAmber500     ColorToken = "primitive.amber.500"
	ColorPrimitiveAmber600     ColorToken = "primitive.amber.600"
	ColorPrimitiveAmber700     ColorToken = "primitive.amber.700"
	ColorPrimitiveAmber800     ColorToken = "primitive.amber.800"
	ColorPrimitiveAmber900     ColorToken = "primitive.amber.900"
	ColorPrimitiveAmber950     ColorToken = "primitive.amber.950"
	ColorPrimitiveYellow50     ColorToken = "primitive.yellow.50"
	ColorPrimitiveYellow100    ColorToken = "primitive.yellow.100"
	ColorPrimitiveYellow200    ColorToken = "primitive.yellow.200"
	ColorPrimitiveYellow300    ColorToken = "primitive.yellow.300"
	ColorPrimitiveYellow400    ColorToken = "primitive.yellow.400"
	ColorPrimitiveYellow500    ColorToken = "primitive.yellow.500"
	ColorPrimitiveYellow600    ColorToken = "primitive.yellow.600"
	ColorPrimitiveYellow700    ColorToken = "primitive.yellow.700"
	ColorPrimitiveYellow800    ColorToken = "primitive.yellow.800"
	ColorPrimitiveYellow900    ColorToken = "primitive.yellow.900"
	ColorPrimitiveYellow950    ColorToken = "primitive.yellow.950"
	ColorPrimitiveLime50       ColorToken = "primitive.lime.50"
	ColorPrimitiveLime100      ColorToken = "primitive.lime.100"
	ColorPrimitiveLime200      ColorToken = "primitive.lime.200"
	ColorPrimitiveLime300      ColorToken = "primitive.lime.300"
	ColorPrimitiveLime400      ColorToken = "primitive.lime.400"
	ColorPrimitiveLime500      ColorToken = "primitive.lime.500"
	ColorPrimitiveLime600      ColorToken = "primitive.lime.600"
	ColorPrimitiveLime700      ColorToken = "primitive.lime.700"
	ColorPrimitiveLime800      ColorToken = "primitive.lime.800"
	ColorPrimitiveLime900      ColorToken = "primitive.lime.900"
	ColorPrimitiveLime950      ColorToken = "primitive.lime.950"
	ColorPrimitiveGreen50      ColorToken = "primitive.green.50"
	ColorPrimitiveGreen100     ColorToken = "primitive.green.100"
	ColorPrimitiveGreen200     ColorToken = "primitive.green.200"
	ColorPrimitiveGreen300     ColorToken = "primitive.green.300"
	ColorPrimitiveGreen400     ColorToken = "primitive.green.400"
	ColorPrimitiveGreen500     ColorToken = "primitive.green.500"
	ColorPrimitiveGreen600     ColorToken = "primitive.green.600"
	ColorPrimitiveGreen700     ColorToken = "primitive.green.700"
	ColorPrimitiveGreen800     ColorToken = "primitive.green.800"
	ColorPrimitiveGreen900     ColorToken = "primitive.green.900"
	ColorPrimitiveGreen950     ColorToken = "primitive.green.950"
	ColorPrimitiveEmerald50    ColorToken = "primitive.emerald.50"
	ColorPrimitiveEmerald100   ColorToken = "primitive.emerald.100"
	ColorPrimitiveEmerald200   ColorToken = "primitive.emerald.200"
	ColorPrimitiveEmerald300   ColorToken = "primitive.emerald.300"
	ColorPrimitiveEmerald400   ColorToken = "primitive.emerald.400"
	ColorPrimitiveEmerald500   ColorToken = "primitive.emerald.500"
	ColorPrimitiveEmerald600   ColorToken = "primitive.emerald.600"
	ColorPrimitiveEmerald700   ColorToken = "primitive.emerald.700"
	ColorPrimitiveEmerald800   ColorToken = "primitive.emerald.800"
	ColorPrimitiveEmerald900   ColorToken = "primitive.emerald.900"
	ColorPrimitiveEmerald950   ColorToken = "primitive.emerald.950"
	ColorPrimitiveTeal50       ColorToken = "primitive.teal.50"
	ColorPrimitiveTeal100      ColorToken = "primitive.teal.100"
	ColorPrimitiveTeal200      ColorToken = "primitive.teal.200"
	ColorPrimitiveTeal300      ColorToken = "primitive.teal.300"
	ColorPrimitiveTeal400      ColorToken = "primitive.teal.400"
	ColorPrimitiveTeal500      ColorToken = "primitive.teal.500"
	ColorPrimitiveTeal600      ColorToken = "primitive.teal.600"
	ColorPrimitiveTeal700      ColorToken = "primitive.teal.700"
	ColorPrimitiveTeal800      ColorToken = "primitive.teal.800"
	ColorPrimitiveTeal900      ColorToken = "primitive.teal.900"
	ColorPrimitiveTeal950      ColorToken = "primitive.teal.950"
	ColorPrimitiveCyan50       ColorToken = "primitive.cyan.50"
	ColorPrimitiveCyan100      ColorToken = "primitive.cyan.100"
	ColorPrimitiveCyan200      ColorToken = "primitive.cyan.200"
	ColorPrimitiveCyan300      ColorToken = "primitive.cyan.300"
	ColorPrimitiveCyan400      ColorToken = "primitive.cyan.400"
	ColorPrimitiveCyan500      ColorToken = "primitive.cyan.500"
	ColorPrimitiveCyan600      ColorToken = "primitive.cyan.600"
	ColorPrimitiveCyan700      ColorToken = "primitive.cyan.700"
	ColorPrimitiveCyan800      ColorToken = "primitive.cyan.800"
	ColorPrimitiveCyan900      ColorToken = "primitive.cyan.900"
	ColorPrimitiveCyan950      ColorToken = "primitive.cyan.950"
	ColorPrimitiveSky50        ColorToken = "primitive.sky.50"
	ColorPrimitiveSky100       ColorToken = "primitive.sky.100"
	ColorPrimitiveSky200       ColorToken = "primitive.sky.200"
	ColorPrimitiveSky300       ColorToken = "primitive.sky.300"
	ColorPrimitiveSky400       ColorToken = "primitive.sky.400"
	ColorPrimitiveSky500       ColorToken = "primitive.sky.500"
	ColorPrimitiveSky600       ColorToken = "primitive.sky.600"
	ColorPrimitiveSky700       ColorToken = "primitive.sky.700"
	ColorPrimitiveSky800       ColorToken = "primitive.sky.800"
	ColorPrimitiveSky900       ColorToken = "primitive.sky.900"
	ColorPrimitiveSky950       ColorToken = "primitive.sky.950"
	ColorPrimitiveBlue50       ColorToken = "primitive.blue.50"
	ColorPrimitiveBlue100      ColorToken = "primitive.blue.100"
	ColorPrimitiveBlue200      ColorToken = "primitive.blue.200"
	ColorPrimitiveBlue300      ColorToken = "primitive.blue.300"
	ColorPrimitiveBlue400      ColorToken = "primitive.blue.400"
	ColorPrimitiveBlue500      ColorToken = "primitive.blue.500"
	ColorPrimitiveBlue600      ColorToken = "primitive.blue.600"
	ColorPrimitiveBlue700      ColorToken = "primitive.blue.700"
	ColorPrimitiveBlue800      ColorToken = "primitive.blue.800"
	ColorPrimitiveBlue900      ColorToken = "primitive.blue.900"
	ColorPrimitiveBlue950      ColorToken = "primitive.blue.950"
	ColorPrimitiveIndigo50     ColorToken = "primitive.indigo.50"
	ColorPrimitiveIndigo100    ColorToken = "primitive.indigo.100"
	ColorPrimitiveIndigo200    ColorToken = "primitive.indigo.200"
	ColorPrimitiveIndigo300    ColorToken = "primitive.indigo.300"
	ColorPrimitiveIndigo400    ColorToken = "primitive.indigo.400"
	ColorPrimitiveIndigo500    ColorToken = "primitive.indigo.500"
	ColorPrimitiveIndigo600    ColorToken = "primitive.indigo.600"
	ColorPrimitiveIndigo700    ColorToken = "primitive.indigo.700"
	ColorPrimitiveIndigo800    ColorToken = "primitive.indigo.800"
	ColorPrimitiveIndigo900    ColorToken = "primitive.indigo.900"
	ColorPrimitiveIndigo950    ColorToken = "primitive.indigo.950"
	ColorPrimitiveViolet50     ColorToken = "primitive.violet.50"
	ColorPrimitiveViolet100    ColorToken = "primitive.violet.100"
	ColorPrimitiveViolet200    ColorToken = "primitive.violet.200"
	ColorPrimitiveViolet300    ColorToken = "primitive.violet.300"
	ColorPrimitiveViolet400    ColorToken = "primitive.violet.400"
	ColorPrimitiveViolet500    ColorToken = "primitive.violet.500"
	ColorPrimitiveViolet600    ColorToken = "primitive.violet.600"
	ColorPrimitiveViolet700    ColorToken = "primitive.violet.700"
	ColorPrimitiveViolet800    ColorToken = "primitive.violet.800"
	ColorPrimitiveViolet900    ColorToken = "primitive.violet.900"
	ColorPrimitiveViolet950    ColorToken = "primitive.violet.950"
	ColorPrimitivePurple50     ColorToken = "primitive.purple.50"
	ColorPrimitivePurple100    ColorToken = "primitive.purple.100"
	ColorPrimitivePurple200    ColorToken = "primitive.purple.200"
	ColorPrimitivePurple300    ColorToken = "primitive.purple.300"
	ColorPrimitivePurple400    ColorToken = "primitive.purple.400"
	ColorPrimitivePurple500    ColorToken = "primitive.purple.500"
	ColorPrimitivePurple600    ColorToken = "primitive.purple.600"
	ColorPrimitivePurple700    ColorToken = "primitive.purple.700"
	ColorPrimitivePurple800    ColorToken = "primitive.purple.800"
	ColorPrimitivePurple900    ColorToken = "primitive.purple.900"
	ColorPrimitivePurple950    ColorToken = "primitive.purple.950"
	ColorPrimitiveFuchsia50    ColorToken = "primitive.fuchsia.50"
	ColorPrimitiveFuchsia100   ColorToken = "primitive.fuchsia.100"
	ColorPrimitiveFuchsia200   ColorToken = "primitive.fuchsia.200"
	ColorPrimitiveFuchsia300   ColorToken = "primitive.fuchsia.300"
	ColorPrimitiveFuchsia400   ColorToken = "primitive.fuchsia.400"
	ColorPrimitiveFuchsia500   ColorToken = "primitive.fuchsia.500"
	ColorPrimitiveFuchsia600   ColorToken = "primitive.fuchsia.600"
	ColorPrimitiveFuchsia700   ColorToken = "primitive.fuchsia.700"
	ColorPrimitiveFuchsia800   ColorToken = "primitive.fuchsia.800"
	ColorPrimitiveFuchsia900   ColorToken = "primitive.fuchsia.900"
	ColorPrimitiveFuchsia950   ColorToken = "primitive.fuchsia.950"
	ColorPrimitivePink50       ColorToken = "primitive.pink.50"
	ColorPrimitivePink100      ColorToken = "primitive.pink.100"
	ColorPrimitivePink200      ColorToken = "primitive.pink.200"
	ColorPrimitivePink300      ColorToken = "primitive.pink.300"
	ColorPrimitivePink400      ColorToken = "primitive.pink.400"
	ColorPrimitivePink500      ColorToken = "primitive.pink.500"
	ColorPrimitivePink600      ColorToken = "primitive.pink.600"
	ColorPrimitivePink700      ColorToken = "primitive.pink.700"
	ColorPrimitivePink800      ColorToken = "primitive.pink.800"
	ColorPrimitivePink900      ColorToken = "primitive.pink.900"
	ColorPrimitivePink950      ColorToken = "primitive.pink.950"
	ColorPrimitiveRose50       ColorToken = "primitive.rose.50"
	ColorPrimitiveRose100      ColorToken = "primitive.rose.100"
	ColorPrimitiveRose200      ColorToken = "primitive.rose.200"
	ColorPrimitiveRose300      ColorToken = "primitive.rose.300"
	ColorPrimitiveRose400      ColorToken = "primitive.rose.400"
	ColorPrimitiveRose500      ColorToken = "primitive.rose.500"
	ColorPrimitiveRose600      ColorToken = "primitive.rose.600"
	ColorPrimitiveRose700      ColorToken = "primitive.rose.700"
	ColorPrimitiveRose800      ColorToken = "primitive.rose.800"
	ColorPrimitiveRose900      ColorToken = "primitive.rose.900"
	ColorPrimitiveRose950      ColorToken = "primitive.rose.950"
	ColorPrimitiveSlate50      ColorToken = "primitive.slate.50"
	ColorPrimitiveSlate100     ColorToken = "primitive.slate.100"
	ColorPrimitiveSlate200     ColorToken = "primitive.slate.200"
	ColorPrimitiveSlate300     ColorToken = "primitive.slate.300"
	ColorPrimitiveSlate400     ColorToken = "primitive.slate.400"
	ColorPrimitiveSlate500     ColorToken = "primitive.slate.500"
	ColorPrimitiveSlate600     ColorToken = "primitive.slate.600"
	ColorPrimitiveSlate700     ColorToken = "primitive.slate.700"
	ColorPrimitiveSlate800     ColorToken = "primitive.slate.800"
	ColorPrimitiveSlate900     ColorToken = "primitive.slate.900"
	ColorPrimitiveSlate950     ColorToken = "primitive.slate.950"
	ColorPrimitiveGray50       ColorToken = "primitive.gray.50"
	ColorPrimitiveGray100      ColorToken = "primitive.gray.100"
	ColorPrimitiveGray200      ColorToken = "primitive.gray.200"
	ColorPrimitiveGray300      ColorToken = "primitive.gray.300"
	ColorPrimitiveGray400      ColorToken = "primitive.gray.400"
	ColorPrimitiveGray500      ColorToken = "primitive.gray.500"
	ColorPrimitiveGray600      ColorToken = "primitive.gray.600"
	ColorPrimitiveGray700      ColorToken = "primitive.gray.700"
	ColorPrimitiveGray800      ColorToken = "primitive.gray.800"
	ColorPrimitiveGray900      ColorToken = "primitive.gray.900"
	ColorPrimitiveGray950      ColorToken = "primitive.gray.950"
	ColorPrimitiveZinc50       ColorToken = "primitive.zinc.50"
	ColorPrimitiveZinc100      ColorToken = "primitive.zinc.100"
	ColorPrimitiveZinc200      ColorToken = "primitive.zinc.200"
	ColorPrimitiveZinc300      ColorToken = "primitive.zinc.300"
	ColorPrimitiveZinc400      ColorToken = "primitive.zinc.400"
	ColorPrimitiveZinc500      ColorToken = "primitive.zinc.500"
	ColorPrimitiveZinc600      ColorToken = "primitive.zinc.600"
	ColorPrimitiveZinc700      ColorToken = "primitive.zinc.700"
	ColorPrimitiveZinc800      ColorToken = "primitive.zinc.800"
	ColorPrimitiveZinc900      ColorToken = "primitive.zinc.900"
	ColorPrimitiveZinc950      ColorToken = "primitive.zinc.950"
	ColorPrimitiveNeutral50    ColorToken = "primitive.neutral.50"
	ColorPrimitiveNeutral100   ColorToken = "primitive.neutral.100"
	ColorPrimitiveNeutral200   ColorToken = "primitive.neutral.200"
	ColorPrimitiveNeutral300   ColorToken = "primitive.neutral.300"
	ColorPrimitiveNeutral400   ColorToken = "primitive.neutral.400"
	ColorPrimitiveNeutral500   ColorToken = "primitive.neutral.500"
	ColorPrimitiveNeutral600   ColorToken = "primitive.neutral.600"
	ColorPrimitiveNeutral700   ColorToken = "primitive.neutral.700"
	ColorPrimitiveNeutral800   ColorToken = "primitive.neutral.800"
	ColorPrimitiveNeutral900   ColorToken = "primitive.neutral.900"
	ColorPrimitiveNeutral950   ColorToken = "primitive.neutral.950"
	ColorPrimitiveStone50      ColorToken = "primitive.stone.50"
	ColorPrimitiveStone100     ColorToken = "primitive.stone.100"
	ColorPrimitiveStone200     ColorToken = "primitive.stone.200"
	ColorPrimitiveStone300     ColorToken = "primitive.stone.300"
	ColorPrimitiveStone400     ColorToken = "primitive.stone.400"
	ColorPrimitiveStone500     ColorToken = "primitive.stone.500"
	ColorPrimitiveStone600     ColorToken = "primitive.stone.600"
	ColorPrimitiveStone700     ColorToken = "primitive.stone.700"
	ColorPrimitiveStone800     ColorToken = "primitive.stone.800"
	ColorPrimitiveStone900     ColorToken = "primitive.stone.900"
	ColorPrimitiveStone950     ColorToken = "primitive.stone.950"
	ColorPrimitiveTaupe50      ColorToken = "primitive.taupe.50"
	ColorPrimitiveTaupe100     ColorToken = "primitive.taupe.100"
	ColorPrimitiveTaupe200     ColorToken = "primitive.taupe.200"
	ColorPrimitiveTaupe300     ColorToken = "primitive.taupe.300"
	ColorPrimitiveTaupe400     ColorToken = "primitive.taupe.400"
	ColorPrimitiveTaupe500     ColorToken = "primitive.taupe.500"
	ColorPrimitiveTaupe600     ColorToken = "primitive.taupe.600"
	ColorPrimitiveTaupe700     ColorToken = "primitive.taupe.700"
	ColorPrimitiveTaupe800     ColorToken = "primitive.taupe.800"
	ColorPrimitiveTaupe900     ColorToken = "primitive.taupe.900"
	ColorPrimitiveTaupe950     ColorToken = "primitive.taupe.950"
	ColorPrimitiveMauve50      ColorToken = "primitive.mauve.50"
	ColorPrimitiveMauve100     ColorToken = "primitive.mauve.100"
	ColorPrimitiveMauve200     ColorToken = "primitive.mauve.200"
	ColorPrimitiveMauve300     ColorToken = "primitive.mauve.300"
	ColorPrimitiveMauve400     ColorToken = "primitive.mauve.400"
	ColorPrimitiveMauve500     ColorToken = "primitive.mauve.500"
	ColorPrimitiveMauve600     ColorToken = "primitive.mauve.600"
	ColorPrimitiveMauve700     ColorToken = "primitive.mauve.700"
	ColorPrimitiveMauve800     ColorToken = "primitive.mauve.800"
	ColorPrimitiveMauve900     ColorToken = "primitive.mauve.900"
	ColorPrimitiveMauve950     ColorToken = "primitive.mauve.950"
	ColorPrimitiveMist50       ColorToken = "primitive.mist.50"
	ColorPrimitiveMist100      ColorToken = "primitive.mist.100"
	ColorPrimitiveMist200      ColorToken = "primitive.mist.200"
	ColorPrimitiveMist300      ColorToken = "primitive.mist.300"
	ColorPrimitiveMist400      ColorToken = "primitive.mist.400"
	ColorPrimitiveMist500      ColorToken = "primitive.mist.500"
	ColorPrimitiveMist600      ColorToken = "primitive.mist.600"
	ColorPrimitiveMist700      ColorToken = "primitive.mist.700"
	ColorPrimitiveMist800      ColorToken = "primitive.mist.800"
	ColorPrimitiveMist900      ColorToken = "primitive.mist.900"
	ColorPrimitiveMist950      ColorToken = "primitive.mist.950"
	ColorPrimitiveOlive50      ColorToken = "primitive.olive.50"
	ColorPrimitiveOlive100     ColorToken = "primitive.olive.100"
	ColorPrimitiveOlive200     ColorToken = "primitive.olive.200"
	ColorPrimitiveOlive300     ColorToken = "primitive.olive.300"
	ColorPrimitiveOlive400     ColorToken = "primitive.olive.400"
	ColorPrimitiveOlive500     ColorToken = "primitive.olive.500"
	ColorPrimitiveOlive600     ColorToken = "primitive.olive.600"
	ColorPrimitiveOlive700     ColorToken = "primitive.olive.700"
	ColorPrimitiveOlive800     ColorToken = "primitive.olive.800"
	ColorPrimitiveOlive900     ColorToken = "primitive.olive.900"
	ColorPrimitiveOlive950     ColorToken = "primitive.olive.950"
)
```

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:86`

## ColorPrimitivePaletteBlack {#api-colorprimitivepaletteblack}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:87`

## ColorPrimitiveRed50 {#api-colorprimitivered50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:88`

## ColorPrimitiveRed100 {#api-colorprimitivered100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:89`

## ColorPrimitiveRed200 {#api-colorprimitivered200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:90`

## ColorPrimitiveRed300 {#api-colorprimitivered300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:91`

## ColorPrimitiveRed400 {#api-colorprimitivered400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:92`

## ColorPrimitiveRed500 {#api-colorprimitivered500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:93`

## ColorPrimitiveRed600 {#api-colorprimitivered600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:94`

## ColorPrimitiveRed700 {#api-colorprimitivered700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:95`

## ColorPrimitiveRed800 {#api-colorprimitivered800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:96`

## ColorPrimitiveRed900 {#api-colorprimitivered900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:97`

## ColorPrimitiveRed950 {#api-colorprimitivered950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:98`

## ColorPrimitiveOrange50 {#api-colorprimitiveorange50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:99`

## ColorPrimitiveOrange100 {#api-colorprimitiveorange100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:100`

## ColorPrimitiveOrange200 {#api-colorprimitiveorange200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:101`

## ColorPrimitiveOrange300 {#api-colorprimitiveorange300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:102`

## ColorPrimitiveOrange400 {#api-colorprimitiveorange400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:103`

## ColorPrimitiveOrange500 {#api-colorprimitiveorange500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:104`

## ColorPrimitiveOrange600 {#api-colorprimitiveorange600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:105`

## ColorPrimitiveOrange700 {#api-colorprimitiveorange700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:106`

## ColorPrimitiveOrange800 {#api-colorprimitiveorange800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:107`

## ColorPrimitiveOrange900 {#api-colorprimitiveorange900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:108`

## ColorPrimitiveOrange950 {#api-colorprimitiveorange950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:109`

## ColorPrimitiveAmber50 {#api-colorprimitiveamber50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:110`

## ColorPrimitiveAmber100 {#api-colorprimitiveamber100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:111`

## ColorPrimitiveAmber200 {#api-colorprimitiveamber200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:112`

## ColorPrimitiveAmber300 {#api-colorprimitiveamber300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:113`

## ColorPrimitiveAmber400 {#api-colorprimitiveamber400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:114`

## ColorPrimitiveAmber500 {#api-colorprimitiveamber500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:115`

## ColorPrimitiveAmber600 {#api-colorprimitiveamber600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:116`

## ColorPrimitiveAmber700 {#api-colorprimitiveamber700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:117`

## ColorPrimitiveAmber800 {#api-colorprimitiveamber800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:118`

## ColorPrimitiveAmber900 {#api-colorprimitiveamber900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:119`

## ColorPrimitiveAmber950 {#api-colorprimitiveamber950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:120`

## ColorPrimitiveYellow50 {#api-colorprimitiveyellow50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:121`

## ColorPrimitiveYellow100 {#api-colorprimitiveyellow100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:122`

## ColorPrimitiveYellow200 {#api-colorprimitiveyellow200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:123`

## ColorPrimitiveYellow300 {#api-colorprimitiveyellow300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:124`

## ColorPrimitiveYellow400 {#api-colorprimitiveyellow400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:125`

## ColorPrimitiveYellow500 {#api-colorprimitiveyellow500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:126`

## ColorPrimitiveYellow600 {#api-colorprimitiveyellow600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:127`

## ColorPrimitiveYellow700 {#api-colorprimitiveyellow700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:128`

## ColorPrimitiveYellow800 {#api-colorprimitiveyellow800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:129`

## ColorPrimitiveYellow900 {#api-colorprimitiveyellow900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:130`

## ColorPrimitiveYellow950 {#api-colorprimitiveyellow950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:131`

## ColorPrimitiveLime50 {#api-colorprimitivelime50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:132`

## ColorPrimitiveLime100 {#api-colorprimitivelime100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:133`

## ColorPrimitiveLime200 {#api-colorprimitivelime200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:134`

## ColorPrimitiveLime300 {#api-colorprimitivelime300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:135`

## ColorPrimitiveLime400 {#api-colorprimitivelime400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:136`

## ColorPrimitiveLime500 {#api-colorprimitivelime500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:137`

## ColorPrimitiveLime600 {#api-colorprimitivelime600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:138`

## ColorPrimitiveLime700 {#api-colorprimitivelime700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:139`

## ColorPrimitiveLime800 {#api-colorprimitivelime800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:140`

## ColorPrimitiveLime900 {#api-colorprimitivelime900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:141`

## ColorPrimitiveLime950 {#api-colorprimitivelime950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:142`

## ColorPrimitiveGreen50 {#api-colorprimitivegreen50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:143`

## ColorPrimitiveGreen100 {#api-colorprimitivegreen100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:144`

## ColorPrimitiveGreen200 {#api-colorprimitivegreen200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:145`

## ColorPrimitiveGreen300 {#api-colorprimitivegreen300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:146`

## ColorPrimitiveGreen400 {#api-colorprimitivegreen400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:147`

## ColorPrimitiveGreen500 {#api-colorprimitivegreen500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:148`

## ColorPrimitiveGreen600 {#api-colorprimitivegreen600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:149`

## ColorPrimitiveGreen700 {#api-colorprimitivegreen700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:150`

## ColorPrimitiveGreen800 {#api-colorprimitivegreen800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:151`

## ColorPrimitiveGreen900 {#api-colorprimitivegreen900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:152`

## ColorPrimitiveGreen950 {#api-colorprimitivegreen950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:153`

## ColorPrimitiveEmerald50 {#api-colorprimitiveemerald50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:154`

## ColorPrimitiveEmerald100 {#api-colorprimitiveemerald100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:155`

## ColorPrimitiveEmerald200 {#api-colorprimitiveemerald200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:156`

## ColorPrimitiveEmerald300 {#api-colorprimitiveemerald300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:157`

## ColorPrimitiveEmerald400 {#api-colorprimitiveemerald400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:158`

## ColorPrimitiveEmerald500 {#api-colorprimitiveemerald500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:159`

## ColorPrimitiveEmerald600 {#api-colorprimitiveemerald600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:160`

## ColorPrimitiveEmerald700 {#api-colorprimitiveemerald700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:161`

## ColorPrimitiveEmerald800 {#api-colorprimitiveemerald800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:162`

## ColorPrimitiveEmerald900 {#api-colorprimitiveemerald900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:163`

## ColorPrimitiveEmerald950 {#api-colorprimitiveemerald950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:164`

## ColorPrimitiveTeal50 {#api-colorprimitiveteal50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:165`

## ColorPrimitiveTeal100 {#api-colorprimitiveteal100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:166`

## ColorPrimitiveTeal200 {#api-colorprimitiveteal200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:167`

## ColorPrimitiveTeal300 {#api-colorprimitiveteal300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:168`

## ColorPrimitiveTeal400 {#api-colorprimitiveteal400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:169`

## ColorPrimitiveTeal500 {#api-colorprimitiveteal500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:170`

## ColorPrimitiveTeal600 {#api-colorprimitiveteal600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:171`

## ColorPrimitiveTeal700 {#api-colorprimitiveteal700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:172`

## ColorPrimitiveTeal800 {#api-colorprimitiveteal800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:173`

## ColorPrimitiveTeal900 {#api-colorprimitiveteal900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:174`

## ColorPrimitiveTeal950 {#api-colorprimitiveteal950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:175`

## ColorPrimitiveCyan50 {#api-colorprimitivecyan50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:176`

## ColorPrimitiveCyan100 {#api-colorprimitivecyan100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:177`

## ColorPrimitiveCyan200 {#api-colorprimitivecyan200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:178`

## ColorPrimitiveCyan300 {#api-colorprimitivecyan300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:179`

## ColorPrimitiveCyan400 {#api-colorprimitivecyan400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:180`

## ColorPrimitiveCyan500 {#api-colorprimitivecyan500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:181`

## ColorPrimitiveCyan600 {#api-colorprimitivecyan600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:182`

## ColorPrimitiveCyan700 {#api-colorprimitivecyan700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:183`

## ColorPrimitiveCyan800 {#api-colorprimitivecyan800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:184`

## ColorPrimitiveCyan900 {#api-colorprimitivecyan900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:185`

## ColorPrimitiveCyan950 {#api-colorprimitivecyan950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:186`

## ColorPrimitiveSky50 {#api-colorprimitivesky50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:187`

## ColorPrimitiveSky100 {#api-colorprimitivesky100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:188`

## ColorPrimitiveSky200 {#api-colorprimitivesky200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:189`

## ColorPrimitiveSky300 {#api-colorprimitivesky300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:190`

## ColorPrimitiveSky400 {#api-colorprimitivesky400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:191`

## ColorPrimitiveSky500 {#api-colorprimitivesky500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:192`

## ColorPrimitiveSky600 {#api-colorprimitivesky600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:193`

## ColorPrimitiveSky700 {#api-colorprimitivesky700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:194`

## ColorPrimitiveSky800 {#api-colorprimitivesky800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:195`

## ColorPrimitiveSky900 {#api-colorprimitivesky900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:196`

## ColorPrimitiveSky950 {#api-colorprimitivesky950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:197`

## ColorPrimitiveBlue50 {#api-colorprimitiveblue50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:198`

## ColorPrimitiveBlue100 {#api-colorprimitiveblue100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:199`

## ColorPrimitiveBlue200 {#api-colorprimitiveblue200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:200`

## ColorPrimitiveBlue300 {#api-colorprimitiveblue300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:201`

## ColorPrimitiveBlue400 {#api-colorprimitiveblue400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:202`

## ColorPrimitiveBlue500 {#api-colorprimitiveblue500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:203`

## ColorPrimitiveBlue600 {#api-colorprimitiveblue600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:204`

## ColorPrimitiveBlue700 {#api-colorprimitiveblue700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:205`

## ColorPrimitiveBlue800 {#api-colorprimitiveblue800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:206`

## ColorPrimitiveBlue900 {#api-colorprimitiveblue900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:207`

## ColorPrimitiveBlue950 {#api-colorprimitiveblue950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:208`

## ColorPrimitiveIndigo50 {#api-colorprimitiveindigo50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:209`

## ColorPrimitiveIndigo100 {#api-colorprimitiveindigo100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:210`

## ColorPrimitiveIndigo200 {#api-colorprimitiveindigo200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:211`

## ColorPrimitiveIndigo300 {#api-colorprimitiveindigo300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:212`

## ColorPrimitiveIndigo400 {#api-colorprimitiveindigo400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:213`

## ColorPrimitiveIndigo500 {#api-colorprimitiveindigo500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:214`

## ColorPrimitiveIndigo600 {#api-colorprimitiveindigo600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:215`

## ColorPrimitiveIndigo700 {#api-colorprimitiveindigo700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:216`

## ColorPrimitiveIndigo800 {#api-colorprimitiveindigo800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:217`

## ColorPrimitiveIndigo900 {#api-colorprimitiveindigo900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:218`

## ColorPrimitiveIndigo950 {#api-colorprimitiveindigo950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:219`

## ColorPrimitiveViolet50 {#api-colorprimitiveviolet50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:220`

## ColorPrimitiveViolet100 {#api-colorprimitiveviolet100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:221`

## ColorPrimitiveViolet200 {#api-colorprimitiveviolet200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:222`

## ColorPrimitiveViolet300 {#api-colorprimitiveviolet300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:223`

## ColorPrimitiveViolet400 {#api-colorprimitiveviolet400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:224`

## ColorPrimitiveViolet500 {#api-colorprimitiveviolet500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:225`

## ColorPrimitiveViolet600 {#api-colorprimitiveviolet600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:226`

## ColorPrimitiveViolet700 {#api-colorprimitiveviolet700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:227`

## ColorPrimitiveViolet800 {#api-colorprimitiveviolet800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:228`

## ColorPrimitiveViolet900 {#api-colorprimitiveviolet900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:229`

## ColorPrimitiveViolet950 {#api-colorprimitiveviolet950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:230`

## ColorPrimitivePurple50 {#api-colorprimitivepurple50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:231`

## ColorPrimitivePurple100 {#api-colorprimitivepurple100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:232`

## ColorPrimitivePurple200 {#api-colorprimitivepurple200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:233`

## ColorPrimitivePurple300 {#api-colorprimitivepurple300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:234`

## ColorPrimitivePurple400 {#api-colorprimitivepurple400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:235`

## ColorPrimitivePurple500 {#api-colorprimitivepurple500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:236`

## ColorPrimitivePurple600 {#api-colorprimitivepurple600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:237`

## ColorPrimitivePurple700 {#api-colorprimitivepurple700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:238`

## ColorPrimitivePurple800 {#api-colorprimitivepurple800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:239`

## ColorPrimitivePurple900 {#api-colorprimitivepurple900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:240`

## ColorPrimitivePurple950 {#api-colorprimitivepurple950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:241`

## ColorPrimitiveFuchsia50 {#api-colorprimitivefuchsia50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:242`

## ColorPrimitiveFuchsia100 {#api-colorprimitivefuchsia100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:243`

## ColorPrimitiveFuchsia200 {#api-colorprimitivefuchsia200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:244`

## ColorPrimitiveFuchsia300 {#api-colorprimitivefuchsia300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:245`

## ColorPrimitiveFuchsia400 {#api-colorprimitivefuchsia400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:246`

## ColorPrimitiveFuchsia500 {#api-colorprimitivefuchsia500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:247`

## ColorPrimitiveFuchsia600 {#api-colorprimitivefuchsia600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:248`

## ColorPrimitiveFuchsia700 {#api-colorprimitivefuchsia700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:249`

## ColorPrimitiveFuchsia800 {#api-colorprimitivefuchsia800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:250`

## ColorPrimitiveFuchsia900 {#api-colorprimitivefuchsia900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:251`

## ColorPrimitiveFuchsia950 {#api-colorprimitivefuchsia950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:252`

## ColorPrimitivePink50 {#api-colorprimitivepink50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:253`

## ColorPrimitivePink100 {#api-colorprimitivepink100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:254`

## ColorPrimitivePink200 {#api-colorprimitivepink200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:255`

## ColorPrimitivePink300 {#api-colorprimitivepink300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:256`

## ColorPrimitivePink400 {#api-colorprimitivepink400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:257`

## ColorPrimitivePink500 {#api-colorprimitivepink500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:258`

## ColorPrimitivePink600 {#api-colorprimitivepink600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:259`

## ColorPrimitivePink700 {#api-colorprimitivepink700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:260`

## ColorPrimitivePink800 {#api-colorprimitivepink800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:261`

## ColorPrimitivePink900 {#api-colorprimitivepink900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:262`

## ColorPrimitivePink950 {#api-colorprimitivepink950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:263`

## ColorPrimitiveRose50 {#api-colorprimitiverose50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:264`

## ColorPrimitiveRose100 {#api-colorprimitiverose100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:265`

## ColorPrimitiveRose200 {#api-colorprimitiverose200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:266`

## ColorPrimitiveRose300 {#api-colorprimitiverose300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:267`

## ColorPrimitiveRose400 {#api-colorprimitiverose400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:268`

## ColorPrimitiveRose500 {#api-colorprimitiverose500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:269`

## ColorPrimitiveRose600 {#api-colorprimitiverose600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:270`

## ColorPrimitiveRose700 {#api-colorprimitiverose700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:271`

## ColorPrimitiveRose800 {#api-colorprimitiverose800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:272`

## ColorPrimitiveRose900 {#api-colorprimitiverose900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:273`

## ColorPrimitiveRose950 {#api-colorprimitiverose950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:274`

## ColorPrimitiveSlate50 {#api-colorprimitiveslate50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:275`

## ColorPrimitiveSlate100 {#api-colorprimitiveslate100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:276`

## ColorPrimitiveSlate200 {#api-colorprimitiveslate200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:277`

## ColorPrimitiveSlate300 {#api-colorprimitiveslate300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:278`

## ColorPrimitiveSlate400 {#api-colorprimitiveslate400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:279`

## ColorPrimitiveSlate500 {#api-colorprimitiveslate500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:280`

## ColorPrimitiveSlate600 {#api-colorprimitiveslate600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:281`

## ColorPrimitiveSlate700 {#api-colorprimitiveslate700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:282`

## ColorPrimitiveSlate800 {#api-colorprimitiveslate800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:283`

## ColorPrimitiveSlate900 {#api-colorprimitiveslate900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:284`

## ColorPrimitiveSlate950 {#api-colorprimitiveslate950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:285`

## ColorPrimitiveGray50 {#api-colorprimitivegray50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:286`

## ColorPrimitiveGray100 {#api-colorprimitivegray100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:287`

## ColorPrimitiveGray200 {#api-colorprimitivegray200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:288`

## ColorPrimitiveGray300 {#api-colorprimitivegray300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:289`

## ColorPrimitiveGray400 {#api-colorprimitivegray400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:290`

## ColorPrimitiveGray500 {#api-colorprimitivegray500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:291`

## ColorPrimitiveGray600 {#api-colorprimitivegray600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:292`

## ColorPrimitiveGray700 {#api-colorprimitivegray700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:293`

## ColorPrimitiveGray800 {#api-colorprimitivegray800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:294`

## ColorPrimitiveGray900 {#api-colorprimitivegray900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:295`

## ColorPrimitiveGray950 {#api-colorprimitivegray950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:296`

## ColorPrimitiveZinc50 {#api-colorprimitivezinc50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:297`

## ColorPrimitiveZinc100 {#api-colorprimitivezinc100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:298`

## ColorPrimitiveZinc200 {#api-colorprimitivezinc200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:299`

## ColorPrimitiveZinc300 {#api-colorprimitivezinc300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:300`

## ColorPrimitiveZinc400 {#api-colorprimitivezinc400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:301`

## ColorPrimitiveZinc500 {#api-colorprimitivezinc500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:302`

## ColorPrimitiveZinc600 {#api-colorprimitivezinc600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:303`

## ColorPrimitiveZinc700 {#api-colorprimitivezinc700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:304`

## ColorPrimitiveZinc800 {#api-colorprimitivezinc800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:305`

## ColorPrimitiveZinc900 {#api-colorprimitivezinc900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:306`

## ColorPrimitiveZinc950 {#api-colorprimitivezinc950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:307`

## ColorPrimitiveNeutral50 {#api-colorprimitiveneutral50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:308`

## ColorPrimitiveNeutral100 {#api-colorprimitiveneutral100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:309`

## ColorPrimitiveNeutral200 {#api-colorprimitiveneutral200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:310`

## ColorPrimitiveNeutral300 {#api-colorprimitiveneutral300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:311`

## ColorPrimitiveNeutral400 {#api-colorprimitiveneutral400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:312`

## ColorPrimitiveNeutral500 {#api-colorprimitiveneutral500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:313`

## ColorPrimitiveNeutral600 {#api-colorprimitiveneutral600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:314`

## ColorPrimitiveNeutral700 {#api-colorprimitiveneutral700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:315`

## ColorPrimitiveNeutral800 {#api-colorprimitiveneutral800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:316`

## ColorPrimitiveNeutral900 {#api-colorprimitiveneutral900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:317`

## ColorPrimitiveNeutral950 {#api-colorprimitiveneutral950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:318`

## ColorPrimitiveStone50 {#api-colorprimitivestone50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:319`

## ColorPrimitiveStone100 {#api-colorprimitivestone100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:320`

## ColorPrimitiveStone200 {#api-colorprimitivestone200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:321`

## ColorPrimitiveStone300 {#api-colorprimitivestone300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:322`

## ColorPrimitiveStone400 {#api-colorprimitivestone400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:323`

## ColorPrimitiveStone500 {#api-colorprimitivestone500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:324`

## ColorPrimitiveStone600 {#api-colorprimitivestone600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:325`

## ColorPrimitiveStone700 {#api-colorprimitivestone700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:326`

## ColorPrimitiveStone800 {#api-colorprimitivestone800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:327`

## ColorPrimitiveStone900 {#api-colorprimitivestone900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:328`

## ColorPrimitiveStone950 {#api-colorprimitivestone950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:329`

## ColorPrimitiveTaupe50 {#api-colorprimitivetaupe50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:330`

## ColorPrimitiveTaupe100 {#api-colorprimitivetaupe100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:331`

## ColorPrimitiveTaupe200 {#api-colorprimitivetaupe200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:332`

## ColorPrimitiveTaupe300 {#api-colorprimitivetaupe300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:333`

## ColorPrimitiveTaupe400 {#api-colorprimitivetaupe400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:334`

## ColorPrimitiveTaupe500 {#api-colorprimitivetaupe500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:335`

## ColorPrimitiveTaupe600 {#api-colorprimitivetaupe600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:336`

## ColorPrimitiveTaupe700 {#api-colorprimitivetaupe700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:337`

## ColorPrimitiveTaupe800 {#api-colorprimitivetaupe800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:338`

## ColorPrimitiveTaupe900 {#api-colorprimitivetaupe900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:339`

## ColorPrimitiveTaupe950 {#api-colorprimitivetaupe950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:340`

## ColorPrimitiveMauve50 {#api-colorprimitivemauve50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:341`

## ColorPrimitiveMauve100 {#api-colorprimitivemauve100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:342`

## ColorPrimitiveMauve200 {#api-colorprimitivemauve200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:343`

## ColorPrimitiveMauve300 {#api-colorprimitivemauve300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:344`

## ColorPrimitiveMauve400 {#api-colorprimitivemauve400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:345`

## ColorPrimitiveMauve500 {#api-colorprimitivemauve500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:346`

## ColorPrimitiveMauve600 {#api-colorprimitivemauve600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:347`

## ColorPrimitiveMauve700 {#api-colorprimitivemauve700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:348`

## ColorPrimitiveMauve800 {#api-colorprimitivemauve800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:349`

## ColorPrimitiveMauve900 {#api-colorprimitivemauve900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:350`

## ColorPrimitiveMauve950 {#api-colorprimitivemauve950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:351`

## ColorPrimitiveMist50 {#api-colorprimitivemist50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:352`

## ColorPrimitiveMist100 {#api-colorprimitivemist100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:353`

## ColorPrimitiveMist200 {#api-colorprimitivemist200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:354`

## ColorPrimitiveMist300 {#api-colorprimitivemist300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:355`

## ColorPrimitiveMist400 {#api-colorprimitivemist400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:356`

## ColorPrimitiveMist500 {#api-colorprimitivemist500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:357`

## ColorPrimitiveMist600 {#api-colorprimitivemist600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:358`

## ColorPrimitiveMist700 {#api-colorprimitivemist700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:359`

## ColorPrimitiveMist800 {#api-colorprimitivemist800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:360`

## ColorPrimitiveMist900 {#api-colorprimitivemist900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:361`

## ColorPrimitiveMist950 {#api-colorprimitivemist950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:362`

## ColorPrimitiveOlive50 {#api-colorprimitiveolive50}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:363`

## ColorPrimitiveOlive100 {#api-colorprimitiveolive100}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:364`

## ColorPrimitiveOlive200 {#api-colorprimitiveolive200}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:365`

## ColorPrimitiveOlive300 {#api-colorprimitiveolive300}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:366`

## ColorPrimitiveOlive400 {#api-colorprimitiveolive400}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:367`

## ColorPrimitiveOlive500 {#api-colorprimitiveolive500}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:368`

## ColorPrimitiveOlive600 {#api-colorprimitiveolive600}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:369`

## ColorPrimitiveOlive700 {#api-colorprimitiveolive700}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:370`

## ColorPrimitiveOlive800 {#api-colorprimitiveolive800}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:371`

## ColorPrimitiveOlive900 {#api-colorprimitiveolive900}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:372`

## ColorPrimitiveOlive950 {#api-colorprimitiveolive950}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `color_tokens.go:373`

## ColorPrimitiveWhite {#api-colorprimitivewhite}


```go
const (
	ColorPrimitiveWhite              ColorToken = "primitive.white"
	ColorPrimitiveBlack              ColorToken = "primitive.black"
	ColorPrimitiveBlue               ColorToken = "primitive.blue"
	ColorPrimitiveGray               ColorToken = "primitive.gray"
	ColorSemanticSurface             ColorToken = "semantic.surface"
	ColorSemanticSurfaceHi           ColorToken = "semantic.surface.high"
	ColorSemanticText                ColorToken = "semantic.text"
	ColorSemanticAccent              ColorToken = "semantic.accent"
	ColorSemanticAccentHover         ColorToken = "semantic.accent.hover"
	ColorSemanticOnAccent            ColorToken = "semantic.on-accent"
	ColorSemanticFocusRing           ColorToken = "semantic.focus-ring"
	ColorSemanticDanger              ColorToken = "semantic.danger"
	ColorSemanticSuccess             ColorToken = "semantic.success"
	ColorSemanticInfo                ColorToken = "semantic.info"
	ColorSemanticWarn                ColorToken = "semantic.warn"
	ColorSemanticBorder              ColorToken = "semantic.border"
	ColorSemanticShadow              ColorToken = "semantic.shadow"
	ColorSemanticScrollTrack         ColorToken = "semantic.scroll-track"
	ColorSemanticScrollThumb         ColorToken = "semantic.scroll-thumb"
	ColorSemanticScrollThumbHover    ColorToken = "semantic.scroll-thumb-hover"
	ColorSemanticScrollThumbActive   ColorToken = "semantic.scroll-thumb-active"
	ColorSemanticScrollThumbDisabled ColorToken = "semantic.scroll-thumb-disabled"
	ColorSemanticSelectHover         ColorToken = "semantic.select-hover"
	ColorSemanticSelectSelected      ColorToken = "semantic.select-selected"
	ColorSemanticSelectDisabled      ColorToken = "semantic.select-disabled"
	ColorSemanticSliderTrack         ColorToken = "semantic.slider-track"
	ColorSemanticSliderFill          ColorToken = "semantic.slider-fill"
	ColorSemanticSliderThumb         ColorToken = "semantic.slider-thumb"
	ColorSemanticSliderThumbBorder   ColorToken = "semantic.slider-thumb-border"
	ColorSemanticProgressTrack       ColorToken = "semantic.progress-track"
	ColorSemanticProgressFill        ColorToken = "semantic.progress-fill"
	ColorSemanticTabsHover           ColorToken = "semantic.tabs-hover"
	ColorSemanticTabsPressed         ColorToken = "semantic.tabs-pressed"
	ColorSemanticTabsSelected        ColorToken = "semantic.tabs-selected"
	ColorSemanticTabsDisabled        ColorToken = "semantic.tabs-disabled"
	ColorSemanticTabsIndicator       ColorToken = "semantic.tabs-indicator"
	ColorSemanticMenuHover           ColorToken = "semantic.menu-hover"
	ColorSemanticMenuActive          ColorToken = "semantic.menu-active"
	ColorSemanticMenuSelected        ColorToken = "semantic.menu-selected"
	ColorSemanticMenuPressed         ColorToken = "semantic.menu-pressed"
	ColorSemanticMenuDisabled        ColorToken = "semantic.menu-disabled"

	MetricPrimitive0                      MetricToken = "primitive.0"
	MetricPrimitive1                      MetricToken = "primitive.1"
	MetricPrimitive2                      MetricToken = "primitive.2"
	MetricSemanticRadius                  MetricToken = "semantic.radius"
	MetricSemanticBorder                  MetricToken = "semantic.border-width"
	MetricSemanticShadow                  MetricToken = "semantic.shadow-blur"
	MetricSemanticTextSize                MetricToken = "semantic.text-size"
	MetricSemanticLineHeight              MetricToken = "semantic.line-height"
	MetricSemanticControlHeight           MetricToken = "semantic.control-height"
	MetricComponentButtonPaddingX         MetricToken = "component.button.padding-x"
	MetricComponentButtonPaddingY         MetricToken = "component.button.padding-y"
	MetricComponentButtonGroupBorderWidth MetricToken = "component.button-group.border-width"
	MetricComponentButtonGroupRadius      MetricToken = "component.button-group.radius"
	MetricComponentBadgePaddingX          MetricToken = "component.badge.padding-x"
	MetricComponentBadgePaddingY          MetricToken = "component.badge.padding-y"
	MetricComponentBadgeMinHeight         MetricToken = "component.badge.min-height"
	MetricComponentBadgeRadius            MetricToken = "component.badge.radius"
	MetricComponentInputPaddingX          MetricToken = "component.input.padding-x"
	MetricComponentInputPaddingY          MetricToken = "component.input.padding-y"
	MetricComponentInputGroupPaddingX     MetricToken = "component.input-group.padding-x"
	MetricComponentInputGroupGap          MetricToken = "component.input-group.gap"
	MetricComponentTextareaMinHeight      MetricToken = "component.textarea.min-height"
	MetricComponentToggleWidth            MetricToken = "component.toggle.width"
	MetricComponentToggleHeight           MetricToken = "component.toggle.height"
	MetricComponentToggleKnobInset        MetricToken = "component.toggle.knob-inset"
	MetricComponentSliderWidth            MetricToken = "component.slider.width"
	MetricComponentSliderHeight           MetricToken = "component.slider.height"
	MetricComponentSliderTrackHeight      MetricToken = "component.slider.track-height"
	MetricComponentSliderThumbSize        MetricToken = "component.slider.thumb-size"
	MetricComponentProgressBarWidth       MetricToken = "component.progress-bar.width"
	MetricComponentProgressBarHeight      MetricToken = "component.progress-bar.height"
	MetricComponentProgressBarTrackHeight MetricToken = "component.progress-bar.track-height"
	MetricComponentProgressBarRadius      MetricToken = "component.progress-bar.radius"
	MetricComponentCheckboxSize           MetricToken = "component.checkbox.size"
	MetricComponentCheckboxGap            MetricToken = "component.checkbox.gap"
	MetricComponentRadioSize              MetricToken = "component.radio.size"
	MetricComponentRadioGap               MetricToken = "component.radio.gap"
	MetricComponentIconSize               MetricToken = "component.icon.size"
	MetricComponentAvatarSize             MetricToken = "component.avatar.size"
	MetricComponentScrollThickness        MetricToken = "component.scroll.thickness"
	MetricComponentScrollMinThumb         MetricToken = "component.scroll.min-thumb"
	MetricComponentScrollInset            MetricToken = "component.scroll.inset"
	MetricComponentSelectPaddingX         MetricToken = "component.select.padding-x"
	MetricComponentSelectPaddingY         MetricToken = "component.select.padding-y"
	MetricComponentSelectItemHeight       MetricToken = "component.select.item-height"
	MetricComponentSelectMaxHeight        MetricToken = "component.select.max-height"
	MetricComponentSelectGap              MetricToken = "component.select.gap"
	MetricComponentPopoverPaddingX        MetricToken = "component.popover.padding-x"
	MetricComponentPopoverPaddingY        MetricToken = "component.popover.padding-y"
	MetricComponentPopoverGap             MetricToken = "component.popover.gap"
	MetricComponentTooltipPaddingX        MetricToken = "component.tooltip.padding-x"
	MetricComponentTooltipPaddingY        MetricToken = "component.tooltip.padding-y"
	MetricComponentTooltipGap             MetricToken = "component.tooltip.gap"
	MetricComponentTabsHeight             MetricToken = "component.tabs.height"
	MetricComponentTabsGap                MetricToken = "component.tabs.gap"
	MetricComponentTabsPaddingX           MetricToken = "component.tabs.padding-x"
	MetricComponentTabsPaddingY           MetricToken = "component.tabs.padding-y"
	MetricComponentTabsIndicator          MetricToken = "component.tabs.indicator-height"
	MetricComponentTabsIndicatorInset     MetricToken = "component.tabs.indicator-inset"
	MetricComponentMenuItemHeight         MetricToken = "component.menu.item-height"
	MetricComponentMenuGap                MetricToken = "component.menu.gap"
	MetricComponentMenuPaddingX           MetricToken = "component.menu.padding-x"
	MetricComponentMenuPaddingY           MetricToken = "component.menu.padding-y"

	ComponentPanel           ComponentToken = "panel"
	ComponentText            ComponentToken = "text"
	ComponentButton          ComponentToken = "button"
	ComponentButtonSecondary ComponentToken = "button.secondary"
	ComponentButtonDanger    ComponentToken = "button.danger"
	ComponentButtonGhost     ComponentToken = "button.ghost"
	ComponentButtonGroup     ComponentToken = "button-group"
	ComponentBadge           ComponentToken = "badge"
	ComponentInput           ComponentToken = "input"
	ComponentInputGroup      ComponentToken = "input-group"
	ComponentToggleSwitch    ComponentToken = "toggle-switch"
	ComponentSlider          ComponentToken = "slider"
	ComponentProgressBar     ComponentToken = "progress-bar"
	ComponentCheckbox        ComponentToken = "checkbox"
	ComponentRadio           ComponentToken = "radio"
	ComponentIcon            ComponentToken = "icon"
	ComponentImage           ComponentToken = "image"
	ComponentAvatar          ComponentToken = "avatar"
	ComponentScroll          ComponentToken = "scroll"
	ComponentSelect          ComponentToken = "select"
	ComponentTabs            ComponentToken = "tabs"
	ComponentMenu            ComponentToken = "menu"
	ComponentPopover         ComponentToken = "popover"
	ComponentTooltip         ComponentToken = "tooltip"
)
```

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:4`

## ColorPrimitiveBlack {#api-colorprimitiveblack}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:5`

## ColorPrimitiveBlue {#api-colorprimitiveblue}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:6`

## ColorPrimitiveGray {#api-colorprimitivegray}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:7`

## ColorSemanticSurface {#api-colorsemanticsurface}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:8`

## ColorSemanticSurfaceHi {#api-colorsemanticsurfacehi}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:9`

## ColorSemanticText {#api-colorsemantictext}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:10`

## ColorSemanticAccent {#api-colorsemanticaccent}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:11`

## ColorSemanticAccentHover {#api-colorsemanticaccenthover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:12`

## ColorSemanticOnAccent {#api-colorsemanticonaccent}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:13`

## ColorSemanticFocusRing {#api-colorsemanticfocusring}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:14`

## ColorSemanticDanger {#api-colorsemanticdanger}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:15`

## ColorSemanticSuccess {#api-colorsemanticsuccess}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:16`

## ColorSemanticInfo {#api-colorsemanticinfo}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:17`

## ColorSemanticWarn {#api-colorsemanticwarn}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:18`

## ColorSemanticBorder {#api-colorsemanticborder}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:19`

## ColorSemanticShadow {#api-colorsemanticshadow}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:20`

## ColorSemanticScrollTrack {#api-colorsemanticscrolltrack}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:21`

## ColorSemanticScrollThumb {#api-colorsemanticscrollthumb}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:22`

## ColorSemanticScrollThumbHover {#api-colorsemanticscrollthumbhover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:23`

## ColorSemanticScrollThumbActive {#api-colorsemanticscrollthumbactive}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:24`

## ColorSemanticScrollThumbDisabled {#api-colorsemanticscrollthumbdisabled}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:25`

## ColorSemanticSelectHover {#api-colorsemanticselecthover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:26`

## ColorSemanticSelectSelected {#api-colorsemanticselectselected}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:27`

## ColorSemanticSelectDisabled {#api-colorsemanticselectdisabled}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:28`

## ColorSemanticSliderTrack {#api-colorsemanticslidertrack}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:29`

## ColorSemanticSliderFill {#api-colorsemanticsliderfill}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:30`

## ColorSemanticSliderThumb {#api-colorsemanticsliderthumb}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:31`

## ColorSemanticSliderThumbBorder {#api-colorsemanticsliderthumbborder}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:32`

## ColorSemanticProgressTrack {#api-colorsemanticprogresstrack}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:33`

## ColorSemanticProgressFill {#api-colorsemanticprogressfill}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:34`

## ColorSemanticTabsHover {#api-colorsemantictabshover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:35`

## ColorSemanticTabsPressed {#api-colorsemantictabspressed}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:36`

## ColorSemanticTabsSelected {#api-colorsemantictabsselected}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:37`

## ColorSemanticTabsDisabled {#api-colorsemantictabsdisabled}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:38`

## ColorSemanticTabsIndicator {#api-colorsemantictabsindicator}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:39`

## ColorSemanticMenuHover {#api-colorsemanticmenuhover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:40`

## ColorSemanticMenuActive {#api-colorsemanticmenuactive}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:41`

## ColorSemanticMenuSelected {#api-colorsemanticmenuselected}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:42`

## ColorSemanticMenuPressed {#api-colorsemanticmenupressed}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:43`

## ColorSemanticMenuDisabled {#api-colorsemanticmenudisabled}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:44`

## MetricPrimitive0 {#api-metricprimitive0}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:46`

## MetricPrimitive1 {#api-metricprimitive1}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:47`

## MetricPrimitive2 {#api-metricprimitive2}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:48`

## MetricSemanticRadius {#api-metricsemanticradius}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:49`

## MetricSemanticBorder {#api-metricsemanticborder}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:50`

## MetricSemanticShadow {#api-metricsemanticshadow}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:51`

## MetricSemanticTextSize {#api-metricsemantictextsize}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:52`

## MetricSemanticLineHeight {#api-metricsemanticlineheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:53`

## MetricSemanticControlHeight {#api-metricsemanticcontrolheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:54`

## MetricComponentButtonPaddingX {#api-metriccomponentbuttonpaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:55`

## MetricComponentButtonPaddingY {#api-metriccomponentbuttonpaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:56`

## MetricComponentButtonGroupBorderWidth {#api-metriccomponentbuttongroupborderwidth}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:57`

## MetricComponentButtonGroupRadius {#api-metriccomponentbuttongroupradius}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:58`

## MetricComponentBadgePaddingX {#api-metriccomponentbadgepaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:59`

## MetricComponentBadgePaddingY {#api-metriccomponentbadgepaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:60`

## MetricComponentBadgeMinHeight {#api-metriccomponentbadgeminheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:61`

## MetricComponentBadgeRadius {#api-metriccomponentbadgeradius}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:62`

## MetricComponentInputPaddingX {#api-metriccomponentinputpaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:63`

## MetricComponentInputPaddingY {#api-metriccomponentinputpaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:64`

## MetricComponentInputGroupPaddingX {#api-metriccomponentinputgrouppaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:65`

## MetricComponentInputGroupGap {#api-metriccomponentinputgroupgap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:66`

## MetricComponentTextareaMinHeight {#api-metriccomponenttextareaminheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:67`

## MetricComponentToggleWidth {#api-metriccomponenttogglewidth}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:68`

## MetricComponentToggleHeight {#api-metriccomponenttoggleheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:69`

## MetricComponentToggleKnobInset {#api-metriccomponenttoggleknobinset}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:70`

## MetricComponentSliderWidth {#api-metriccomponentsliderwidth}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:71`

## MetricComponentSliderHeight {#api-metriccomponentsliderheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:72`

## MetricComponentSliderTrackHeight {#api-metriccomponentslidertrackheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:73`

## MetricComponentSliderThumbSize {#api-metriccomponentsliderthumbsize}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:74`

## MetricComponentProgressBarWidth {#api-metriccomponentprogressbarwidth}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:75`

## MetricComponentProgressBarHeight {#api-metriccomponentprogressbarheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:76`

## MetricComponentProgressBarTrackHeight {#api-metriccomponentprogressbartrackheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:77`

## MetricComponentProgressBarRadius {#api-metriccomponentprogressbarradius}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:78`

## MetricComponentCheckboxSize {#api-metriccomponentcheckboxsize}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:79`

## MetricComponentCheckboxGap {#api-metriccomponentcheckboxgap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:80`

## MetricComponentRadioSize {#api-metriccomponentradiosize}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:81`

## MetricComponentRadioGap {#api-metriccomponentradiogap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:82`

## MetricComponentIconSize {#api-metriccomponenticonsize}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:83`

## MetricComponentAvatarSize {#api-metriccomponentavatarsize}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:84`

## MetricComponentScrollThickness {#api-metriccomponentscrollthickness}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:85`

## MetricComponentScrollMinThumb {#api-metriccomponentscrollminthumb}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:86`

## MetricComponentScrollInset {#api-metriccomponentscrollinset}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:87`

## MetricComponentSelectPaddingX {#api-metriccomponentselectpaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:88`

## MetricComponentSelectPaddingY {#api-metriccomponentselectpaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:89`

## MetricComponentSelectItemHeight {#api-metriccomponentselectitemheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:90`

## MetricComponentSelectMaxHeight {#api-metriccomponentselectmaxheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:91`

## MetricComponentSelectGap {#api-metriccomponentselectgap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:92`

## MetricComponentPopoverPaddingX {#api-metriccomponentpopoverpaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:93`

## MetricComponentPopoverPaddingY {#api-metriccomponentpopoverpaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:94`

## MetricComponentPopoverGap {#api-metriccomponentpopovergap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:95`

## MetricComponentTooltipPaddingX {#api-metriccomponenttooltippaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:96`

## MetricComponentTooltipPaddingY {#api-metriccomponenttooltippaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:97`

## MetricComponentTooltipGap {#api-metriccomponenttooltipgap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:98`

## MetricComponentTabsHeight {#api-metriccomponenttabsheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:99`

## MetricComponentTabsGap {#api-metriccomponenttabsgap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:100`

## MetricComponentTabsPaddingX {#api-metriccomponenttabspaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:101`

## MetricComponentTabsPaddingY {#api-metriccomponenttabspaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:102`

## MetricComponentTabsIndicator {#api-metriccomponenttabsindicator}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:103`

## MetricComponentTabsIndicatorInset {#api-metriccomponenttabsindicatorinset}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:104`

## MetricComponentMenuItemHeight {#api-metriccomponentmenuitemheight}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:105`

## MetricComponentMenuGap {#api-metriccomponentmenugap}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:106`

## MetricComponentMenuPaddingX {#api-metriccomponentmenupaddingx}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:107`

## MetricComponentMenuPaddingY {#api-metriccomponentmenupaddingy}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:108`

## ComponentPanel {#api-componentpanel}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:110`

## ComponentText {#api-componenttext}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:111`

## ComponentButton {#api-componentbutton}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:112`

## ComponentButtonSecondary {#api-componentbuttonsecondary}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:113`

## ComponentButtonDanger {#api-componentbuttondanger}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:114`

## ComponentButtonGhost {#api-componentbuttonghost}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:115`

## ComponentButtonGroup {#api-componentbuttongroup}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:116`

## ComponentBadge {#api-componentbadge}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:117`

## ComponentInput {#api-componentinput}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:118`

## ComponentInputGroup {#api-componentinputgroup}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:119`

## ComponentToggleSwitch {#api-componenttoggleswitch}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:120`

## ComponentSlider {#api-componentslider}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:121`

## ComponentProgressBar {#api-componentprogressbar}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:122`

## ComponentCheckbox {#api-componentcheckbox}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:123`

## ComponentRadio {#api-componentradio}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:124`

## ComponentIcon {#api-componenticon}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:125`

## ComponentImage {#api-componentimage}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:126`

## ComponentAvatar {#api-componentavatar}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:127`

## ComponentScroll {#api-componentscroll}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:128`

## ComponentSelect {#api-componentselect}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:129`

## ComponentTabs {#api-componenttabs}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:130`

## ComponentMenu {#api-componentmenu}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:131`

## ComponentPopover {#api-componentpopover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:132`

## ComponentTooltip {#api-componenttooltip}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:133`

## LightTheme {#api-lighttheme}

```text
LightTheme returns an independent, mutable-by-the-caller light theme value.
App.SetTheme copies it before use.
```


```go
func LightTheme() Theme
```

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:138`

## DarkTheme {#api-darktheme}

```text
DarkTheme returns an independent, mutable-by-the-caller dark theme value.
```


```go
func DarkTheme() Theme
```

[对应示例](/zh-cn/api/theme) · 源文件 `theme.go:143`

