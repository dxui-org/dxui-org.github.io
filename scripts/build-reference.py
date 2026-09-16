"""Rebuild source-backed reference pages; run api-scan.go first."""
import json
import re
from pathlib import Path
from i18n import english, chinese

ROOT = Path(__file__).resolve().parents[1]
DOCS = ROOT / 'docs'
API = json.loads((ROOT / 'scripts/api-snapshot.json').read_text(encoding='utf-8-sig'))
# The source audit stays complete; the website omits backend implementation details.
for entry in API:
    entry['Doc'] = entry['Doc'].replace('backend or SDL handles', 'backend handles')
    entry['Doc'] = entry['Doc'].replace('without loading SDL', 'without starting the native runtime')
    entry['Doc'] = entry['Doc'].replace('SDL-free', 'backend-independent')
    if entry['Name'] == 'WindowOptions':
        entry['Doc'] = 'WindowOptions configures an independent child window. Fonts, theme, rendering and cache settings come from the owner App. Shortcuts are configured separately for each window.'
    entry['Code'] = '\n'.join(line for line in entry['Code'].splitlines() if 'SDLVersion' not in line)
def write(path, value):
    if re.search(r'\bsdl\w*', value, re.IGNORECASE):
        raise ValueError(f'{path}: backend implementation details must not appear on the website')
    p = DOCS / path
    p.parent.mkdir(parents=True, exist_ok=True)
    if path.endswith('.md'):
        en = english(value)
        zh = DOCS / 'zh-cn' / path
        zh.parent.mkdir(parents=True, exist_ok=True)
        zh.write_text(chinese(path, value), encoding='utf-8')
        p.write_text(en, encoding='utf-8')
    else:
        p.write_text(value, encoding='utf-8')
def slug(name):
    return re.sub(r'(?<!^)(?=[A-Z])', '-', name).lower()
def code(s): return '\n```go\n' + s.strip() + '\n```\n'

# Each example is a complete program; state lives outside the root builder.
COMPONENTS = {
'Box': ('容器与布局', '', 'dxui.Box(dxui.BoxProps{Direction: dxui.Horizontal, Gap: 12, Align: dxui.AlignCenter}, dxui.Label("Left"), dxui.TextButton(dxui.ButtonProps{}, "Right"))', '默认 Vertical，单行排列；Gap 为逻辑单位。Justify 控制主轴，Align 控制交叉轴。children 可为空，但不能包含零 View。不支持换行、order、反向排列或完整 CSS Flexbox。'),
'Text': ('文本与 Label', '', 'dxui.Text(dxui.TextProps{Value: "Hello dxui", Wrap: dxui.TextWrapWords, MaxLines: 2, Style: dxui.Style{Width: dxui.Px(240), Text: dxui.TextStyle{Size: 20}}})', 'Label(value) 等价于默认 Text；TextNoWrap 为零值。TextWrapWords 使用简单分词换行，合并空白，并非完整 Unicode 换行。MaxLines 为正时限制行数。无效 UTF-8 替换为 U+FFFD。文本本身不是可选择的编辑器；中文依赖字体覆盖。'),
'Button': ('按钮与 TextButton', 'count := 0', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Label(fmt.Sprint(count)), dxui.TextButton(dxui.ButtonProps{Variant: dxui.ButtonOutline, Tone: dxui.ButtonSuccess, Size: dxui.ButtonNormal, OnPress: func(){ count++ }}, "Add"), dxui.Button(dxui.ButtonProps{Disabled: true}, dxui.Label("Disabled")))', 'Button 接受一个任意子 View；TextButton 自动居中显示文字。默认 Filled / Primary / Normal。Variant 支持 Filled、Soft、Outline、Dashed、Ghost、Link；Tone 支持 Primary、Secondary、Success、Info、Warn、Danger；Size 支持 Small、Normal、Large。显式 Token 替换外观配方，局部 Style 保持覆盖权。指针或 Enter/Space 匹配释放时激活一次；nil OnPress 仍保留焦点与交互外观。Link 不执行导航。'),
'ButtonGroup': ('连接按钮组', 'choice := "A"', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Label(choice), dxui.ButtonGroup(dxui.ButtonGroupProps{Dividers: true}, dxui.TextButton(dxui.ButtonProps{OnPress: func(){choice="A"}}, "A"), dxui.TextButton(dxui.ButtonProps{OnPress: func(){choice="B"}}, "B")))', '仅接收 Button（含 TextButton），允许零个。默认水平且没有分隔线；Orientation=ButtonGroupVertical 改为纵向。Dividers 启用主题边框重叠，使共享边只绘制一次。组无独立焦点，每个按钮保留 Disabled、回调和源顺序 Tab；没有方向键导航。纵向自动宽度对齐最宽项，显式子 Width 优先。'),
'InputGroup': ('带前后缀的输入框', 'value := ""', 'dxui.InputGroup(dxui.InputGroupProps{Style: dxui.Style{Width: dxui.Px(360)}}, dxui.InputGroupContent{Input: dxui.Input(dxui.InputProps{Key: "query", Value: value, OnChange: dxui.Assign(&value)}), Prefix: dxui.Some(dxui.Label("Search")), Suffix: dxui.Some(dxui.TextButton(dxui.ButtonProps{OnPress: func(){value=""}}, "Clear"))})', 'InputGroupContent.Input 必须恰好是一个 Input；Prefix/Suffix 用 Some 显式提供，可包含非交互内容或 Button，不能包含其它可聚焦控件或嵌套编辑器。组统一绘制背景、边框、圆角和 focus-within，Input 使用剩余宽度；编辑器状态与 Key 保留。后缀按钮是独立 Tab 停靠点。'),
'Input': ('单行输入', 'value := ""\nselection := dxui.TextRange{}\nmessage := "Press Enter"', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Input(dxui.InputProps{Key: "password", Value: value, OnChange: dxui.Assign(&value), Selection: dxui.Some(selection), OnSelectionChange: dxui.Assign(&selection), Password: true, ShowPasswordToggle: true, Placeholder: "Password", OnSubmit: func(){message="Submitted"}}), dxui.Label(message))', 'Value 是完整受控字符串；不写回 OnChange 提案就拒绝编辑。Selection 用 rune 下标，Some 后必须接纳 OnSelectionChange 才能移动选区。ReadOnly 保留导航、选区、非密码复制与 OnSubmit；Disabled 移除焦点并停止输入。密码即使显示明文也禁止复制/剪切。ShowPasswordToggle 仅在 Password=true 时有效，只改变内部显示，不触发 OnChange。IME 组合阶段不触发 OnChange；提交后才产生值。'),
'Textarea': ('多行输入', 'value := "First line\\nSecond line"', 'dxui.Textarea(dxui.TextareaProps{Key: "notes", Value: value, OnChange: dxui.Assign(&value), Wrap: dxui.TextWrapWords, Placeholder: "Notes", Style: dxui.Style{Width: dxui.Px(400), Height: dxui.Px(180)}})', '遵循 Input 的受控值与 rune 选区规则。Enter 在可编辑状态插入换行，无 OnSubmit、Password 或密码按钮。Wrap 默认 TextNoWrap，可选 TextWrapWords。保留内部光标滚动和撤销历史；ReadOnly 仍允许选择、复制和滚动。不是富文本编辑器，不支持 grapheme 编辑或 RTL 排版。'),
'Checkbox': ('复选框', 'checked := false', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Checkbox(dxui.CheckboxProps{Checked: checked, OnChange: dxui.Assign(&checked)}, dxui.Label("Accept terms")), dxui.Label(fmt.Sprint(checked)))', '接收一个标签 View。Checked 为受控布尔值，鼠标或 Space 提案取反；nil OnChange 不改变业务值，不等同 Disabled。Disabled 移除焦点并取消按压。无三态。'),
'Radio': ('单选按钮', 'choice := "small"', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Radio(dxui.RadioProps{Selected: choice=="small", OnSelect: func(){choice="small"}}, dxui.Label("Small")), dxui.Radio(dxui.RadioProps{Selected: choice=="large", OnSelect: func(){choice="large"}}, dxui.Label("Large")), dxui.Label(choice))', 'Selected 是权威值。启用且未选中时，指针或 Space 调用 OnSelect；已选中不会反选，也不回调。互斥由应用共享业务变量实现，没有 RadioGroup。每个 Radio 保留自己的焦点停靠点。'),
'ToggleSwitch': ('开关', 'checked := false', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.ToggleSwitch(dxui.ToggleSwitchProps{Checked: checked, OnChange: dxui.Assign(&checked)}), dxui.Label(fmt.Sprint(checked)))', 'Checked / OnChange 与 Checkbox 相同，通过指针或 Space 提案。默认尺寸由主题指标决定（40×24 逻辑单位）。没有内置文字标签，可与 Text/Label 组合；nil 回调仍有交互外观。'),
'Slider': ('滑块', 'value := float32(25)', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Slider(dxui.SliderProps{Min: 0, Max: 100, Step: 5, Value: value, OnChange: dxui.Assign(&value)}), dxui.Label(fmt.Sprint(value)))', 'Min 和 Max 同为零表示 0..100；Step 非正或非有限时回退 1。其它边界按字面使用，非有限边界或 Min>=Max 使控件不可交互、不可聚焦。Value 超范围被钳制；NaN/-Inf 显示 Min，+Inf 显示 Max。提案以 Min 为起点对齐步长，端点仍可到达。支持轨道点击、捕获拖动、方向键、Home/End。'),
'Select': ('下拉选择', 'value := "go"', 'dxui.Select(dxui.SelectProps{Value: value, OnChange: dxui.Assign(&value), Placeholder: "Choose", Options: []dxui.SelectOption{{Value:"go", Label:"Go"}, {Value:"rust", Label:"Rust"}, {Value:"later", Label:"Later", Disabled:true}}})', 'Options 的 Value 必须非空且唯一，作为重排身份。支持空选项列表和未匹配 Value。nil OnChange 仍可打开、浏览、关闭，只不接纳选择。支持 Enter/Space、方向键、Home/End、Escape、Tab；禁用选项跳过。弹层逃逸祖先 Scroll 裁剪，并在窗口边缘翻转/钳制。不是原生系统选择框，不虚拟化，最多 4096 项。'),
'Tabs': ('标签页选择器', 'value := "home"', 'dxui.Box(dxui.BoxProps{Gap: 12}, dxui.Tabs(dxui.TabsProps{Value:value, OnChange:dxui.Assign(&value), Items:[]dxui.TabItem{{Value:"home",Label:"Home"},{Value:"settings",Label:"Settings"}}}), dxui.Label("Panel: "+value))', '只渲染横向标签，内容面板由应用依据 Value 构建。Items.Value 必须非空且唯一；空/未匹配 Value 不显示选中指示。左右/Home/End 移动内部当前项，Enter/Space 才提案；已选中或 Disabled 项不回调。没有 TabPanel、垂直、滚动、关闭或拖拽排序 API。'),
'Menu': ('菜单动作列表', 'value := ""\nactions := 0', 'dxui.Box(dxui.BoxProps{Gap:12}, dxui.Menu(dxui.MenuProps{Value:value, Items:[]dxui.MenuItem{{Value:"save",Label:"Save"},{Value:"export",Label:"Export"}}, OnAction:func(v string){value=v; actions++}}), dxui.Label(fmt.Sprintf("%s: %d actions", value, actions)))', '普通布局内的动作列表，默认 MenuVertical，MenuHorizontal 改为水平。Value 是可选受控选择外观；OnAction 对已选中项仍会再次触发。Item.Value 非空唯一。一个菜单贡献一个 Tab 停靠点；对应方向键/Home/End 移动内部当前项，不循环，Enter/Space 执行。无内置弹层、子菜单、分隔线、路由或快捷键字段。长菜单组合 Scroll。'),
'Scroll': ('滚动容器', 'offset := dxui.Point{}', 'dxui.Scroll(dxui.ScrollProps{Style:dxui.Style{Height:dxui.Px(240), Width:dxui.Px(400)}, Offset:dxui.Some(offset), OnScroll:dxui.Assign(&offset), Scrollbar:dxui.ScrollbarAlways}, dxui.Box(dxui.BoxProps{Gap:12}, dxui.Label("Top"), dxui.Box(dxui.BoxProps{Style:dxui.Style{Height:dxui.Px(600), Shrink:dxui.NoShrink()}}), dxui.Label("Bottom")))', '恰好一个 child。Axis 默认 ScrollVertical，可选 ScrollHorizontal/ScrollBoth；启用轴对子项无上限测量，因此视口需要确定大小。Offset 未设置时保留内部偏移，InitialOffset 只首次挂载使用；Some Offset 后 OnScroll 仅提案。ScrollbarHidden 仍允许滚轮。嵌套滚动在边界传递剩余量。没有惯性/平滑滚动；普通 Scroll 挂载全部内容。'),
'VirtualList': ('固定行高虚拟列表', 'items := make([]string, 10000)\nfor i := range items { items[i] = fmt.Sprintf("item-%d",i) }', 'dxui.VirtualList(dxui.VirtualListProps{Key:"items", Style:dxui.Style{Height:dxui.Px(320),Width:dxui.Px(400)}, Count:len(items), Version:1, RowHeight:36, Overscan:3, ItemKey:func(i int)string{return items[i]}, Build:func(i int)dxui.View{return dxui.Label(items[i])}})', '仅垂直固定行高；必须给出有限的像素 Height。RowHeight 包括间距，行被限制为该高度并裁剪。Count 上限 MaxVirtualListItems=10,000,000，Overscan 上限 MaxVirtualListOverscan=256。ItemKey 全列表非空唯一；Count/Version 是不可变快照，内容或键变更必须递增 Version。回调在 UI 线程且应纯粹。离开 overscan 即卸载，焦点、IME、编辑和弹层状态不恢复；业务值放在列表外。非受控滚动按首个幸存可见键锚定；受控 Offset 优先。'),
'Popover': ('交互弹层', 'open := false', 'dxui.Popover(dxui.PopoverProps{Open:open, OnOpenChange:dxui.Assign(&open), Placement:dxui.OverlayBottomStart, Offset:dxui.Metric(6)}, dxui.Label("Open details"), dxui.Box(dxui.BoxProps{Gap:12}, dxui.Label("Details"), dxui.TextButton(dxui.ButtonProps{OnPress:func(){open=false}},"Close")))', '接受 anchor、content 两个 View；Open 必须由应用接纳 OnOpenChange。触发器指针/Enter/Space、顶层 Escape 或外部主键点击会提案。弹层位于窗口级，自动翻转/钳制；Tab 可转入内容，关闭时恢复仍存活触发器焦点。没有 Disabled 字段；禁用 anchor 子组件不等于禁用 Popover 宿主。无模态、焦点陷阱、箭头或动画。'),
'Tooltip': ('提示浮层', '', 'dxui.Tooltip(dxui.TooltipProps{Placement:dxui.OverlayTop, Delay:700*time.Millisecond}, dxui.TextButton(dxui.ButtonProps{},"Hover or Tab here"), dxui.Label("A helpful hint"))', '接受 anchor、content；鼠标悬停或键盘焦点启动 Delay，离开且失焦后关闭。Delay=0 表示默认 500ms，不是立即出现。Disabled 抑制提示。内容不可交互，不抢焦点、不阻挡指针；需要按钮等交互内容时使用 Popover。由事件/截止时间驱动，无需应用计时循环。'),
'Icon': ('矢量图标', '', 'dxui.Icon(dxui.IconProps{Data:icon.Search(), Size:32, StrokeWidth:2, Color:dxui.TokenColor(dxui.Color.Semantic.Accent)})', 'Size 是逻辑单位，0 使用主题默认 20。StrokeWidth 为 view-box 单位，0 表示默认 2，必须有限且非负；对自定义填充图标忽略。图标无语义动作，需要组合 Button。icon.Name() 按需链接；icon/catalog 会引入完整目录。IconData 不是 SVG 解析器。'),
'Image': ('栅格图片', 'pixels := image.NewRGBA(image.Rect(0,0,64,64))\nfor y:=0;y<64;y++ { for x:=0;x<64;x++ {pixels.SetRGBA(x,y,color.RGBA{R:uint8(x*4),G:uint8(y*4),B:160,A:255})} }\nsource := dxui.ImageFromGo(pixels)\nstatus := "Loading"', 'dxui.Box(dxui.BoxProps{Gap:12}, dxui.Image(dxui.ImageProps{Source:source, Fit:dxui.ImageContain, Alignment:dxui.Point{X:.5,Y:.5}, MaxPixels:1024*1024, Style:dxui.Style{Width:dxui.Px(180),Height:dxui.Px(120)}, OnLoad:func(s dxui.Size){status=fmt.Sprint(s)}, OnError:func(err error){status=err.Error()}}),dxui.Label(status))', 'Source 由 ImageBytes、ImageFile 或 ImageFromGo 构造；放在 builder 外复用。Fit 默认 Contain（保比例完整显示），Cover 填满并裁剪，Fill 拉伸，None 原始尺寸。Alignment 每轴 0..1，默认左上。MaxPixels 限制解码像素数；只有 PNG/JPEG/单帧 GIF，无网络 URL 或异步解码。nil 加载回调不阻止解码。Go 图像构造后不得修改。'),
'Avatar': ('头像', 'pixels := image.NewRGBA(image.Rect(0,0,48,48))\nfor y:=0;y<48;y++ { for x:=0;x<48;x++ {pixels.SetRGBA(x,y,color.RGBA{R:40,G:160,B:120,A:255})} }\nsource := dxui.ImageFromGo(pixels)', 'dxui.Avatar(dxui.AvatarProps{Source:source, Shape:dxui.AvatarCircle, Size:64})', '共享 ImageSource 和 OnLoad/OnError，固定居中 Cover。Shape 默认 AvatarCircle，可选 AvatarSquare。Size=0 用主题默认 40，正数为逻辑尺寸；显式 Style 宽高优先。没有 URL、文字缩写、自动降级或状态徽点；需要降级时由应用根据 OnError 改渲染内容。'),
'Badge': ('徽章标签', '', 'dxui.Badge(dxui.BadgeProps{Style:dxui.Style{Background:dxui.TokenColor(dxui.Color.Semantic.Success)}},dxui.Label("Ready"))', '接受一个任意子 View，继承文字/图标色。默认紧凑胶囊形；无事件、焦点或保留交互状态。不内置通知角标定位、数字封顶、关闭、状态颜色变体或动画，应用用布局和 Style 组合。'),
'ProgressBar': ('进度条', 'value := float32(.25)', 'dxui.Box(dxui.BoxProps{Gap:12},dxui.ProgressBar(dxui.ProgressBarProps{Value:value}),dxui.TextButton(dxui.ButtonProps{OnPress:func(){value+=.1;if value>1{value=0}}},"Advance"),dxui.Label(fmt.Sprintf("%.0f%%",value*100)))', 'Value 为 0..1 比例，不是百分数；越界钳制，NaN/-Inf 为空，+Inf 为满。无事件和内部进度计时。Background 改轨道色，Text.Color 改填充色；默认 160×12，轨道高 8。没有不确定进度、分段、竖向或动画模式。'),
}

FIELDS = {'Min':'下界；与 Max 同为 0 时使用默认 0..100，否则按字面值；须有限且小于 Max。', 'Max':'上界；与 Min 同为 0 时默认 100，否则按字面值；须有限且大于 Min。', 'Step':'以 Min 为起点的步长；非正或非有限时默认 1，端点仍可到达。', 'Key':'同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。', 'Style':'局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。', 'Token':'组件主题条目；空值选择该组件默认。', 'States':'真实交互状态的绘制补丁，不能伪造状态或改变布局。', 'Pointer':'默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。', 'Value':'应用拥有的当前值；行为、范围和零值见本页说明。', 'Disabled':'默认 false；交互组件为 true 时移除焦点并取消交互。', 'Checked':'当前受控布尔值，默认 false；在 OnChange 中接纳提案。', 'Selected':'当前是否选中，默认 false；与共享业务值比较得到。', 'OnPress':'激活后在 UI 线程执行；nil 不执行动作。', 'OnChange':'完整新值提案；写回业务变量后下一次构建生效。nil 行为见本页。', 'OnSelect':'未选中 Radio 激活时调用，无参数；由应用更新共享选择值。', 'OnAction':'每次执行启用菜单项时接收 Value；允许重复执行同一项。', 'Direction':'默认 Vertical；Horizontal 从左到右排列。', 'Gap':'相邻子项的非负固定逻辑间距，默认 0。', 'Justify':'主轴 Start/Center/End/SpaceBetween，默认 Start。', 'Align':'交叉轴 Start/Center/End/Stretch，默认 Start。', 'Selection':'默认未设置，使用内部选区；Some(TextRange) 后为受控 rune 选区。', 'OnSelectionChange':'完整 rune 选区提案；nil 保留内部选区，受控 Selection 在重建时优先。', 'Placeholder':'空 Value 时显示；默认空字符串。', 'ReadOnly':'默认 false；true 禁止编辑和预编辑，保留导航与选择。', 'Password':'默认 false；true 遮蔽显示并禁止复制/剪切。', 'ShowPasswordToggle':'默认 false；仅 Password=true 生效，保留身份期间记住显示状态。', 'OnSubmit':'Input 的 Enter 回调；ReadOnly 时仍可调用，Disabled 时不可。', 'Wrap':'默认 TextNoWrap；TextWrapWords 简单按词换行。', 'MaxLines':'正数限制可见行数；0 不限制，负数使构建失败。', 'Variant':'默认 ButtonFilled；外观配方，详见本页。', 'Tone':'默认 ButtonPrimary；ButtonDefault 是其别名。', 'Size':'尺寸；Button 是枚举，其它组件为逻辑单位，0 取主题默认。', 'Orientation':'默认水平（ButtonGroup）或垂直（Menu）；详见本页。', 'Dividers':'默认 false；true 绘制组内主题分隔边。', 'Axis':'默认 ScrollVertical；启用轴无界测量子项。', 'InitialOffset':'可选首次挂载偏移；后续修改不会重置内部位置。', 'Offset':'Scroll/VirtualList 为可选受控 Point；浮层为可选语义的 MetricValue 间距。', 'Scrollbar':'默认 ScrollbarAuto；Always 始终显示，Hidden 隐藏轨道但仍可滚动。', 'OnScroll':'完整 Point 提案；未设置 Offset 时内部滚动不依赖此回调。', 'Count':'快照条目数；0 为空，最大 MaxVirtualListItems。', 'Version':'快照版本 uint64；内容/键改变时递增，不能在 Build 内修改。', 'RowHeight':'每行完整有限正逻辑高度，包括间距。', 'Overscan':'视口外保留的行数，0..MaxVirtualListOverscan。', 'ItemKey':'按索引返回全列表稳定、非空、唯一键；UI 线程纯函数。', 'Build':'按索引构建行 View；只为挂载窗口构建，不可有副作用。', 'Options':'复制的 SelectOption 列表；值非空唯一，上限 4096。', 'Items':'复制的 TabItem/MenuItem 列表；值非空唯一，允许空列表。', 'Open':'受控打开状态，默认 false。', 'OnOpenChange':'打开/关闭布尔提案；nil 不改变 Open。', 'Placement':'默认 OverlayBottomStart；共八个方向/对齐枚举，见参考。', 'Delay':'time.Duration；0 使用 500ms 默认延迟，负数无效。', 'Data':'有效 IconData；可用官方 icon 构造器或自定义填充路径。', 'Color':'图标显式颜色或 TokenColor；未设置时继承主题色。', 'StrokeWidth':'有限非负 view-box 描边宽，0 为 2；填充路径忽略。', 'Source':'稳定 ImageSource；复用资源句柄，避免每次构建重新分配。', 'Fit':'默认 ImageContain；支持 Cover、Fill、None。', 'Alignment':'X/Y 均为 0..1；默认 (0,0) 左上，(.5,.5) 居中。', 'MaxPixels':'解码像素预算；0 选引擎默认 16,777,216，负值无效。', 'OnLoad':'加载成功通知，参数为解码 Size；nil 仍加载。', 'OnError':'资源错误通知；应用可切换占位视图，nil 仍尝试加载。', 'Shape':'默认 AvatarCircle；AvatarSquare 为方形。'}

def program(name, state, expr):
    imports=['"log"','"github.com/dxui-org/dxui"']
    body=state+'\n'+expr
    for key, imp in [('fmt.','fmt'),('time.','time'),('image.','image'),('color.','image/color'),('icon.','github.com/dxui-org/dxui/icon')]:
        if key in body: imports.append('"'+imp+'"')
    return 'package main\n\nimport (\n'+ '\n'.join(imports)+'\n)\n\nfunc main() {\n'+state+'\napp := dxui.NewApp(dxui.AppOptions{Title:"'+name+'", Width:640, Height:480, Background:dxui.RGBA(248,250,252,255)})\nif err := app.Run(func() dxui.View {\nreturn dxui.Box(dxui.BoxProps{Gap:16, Style:dxui.Style{Padding:dxui.Padding(24)}},\n'+expr+',\n)\n}); err != nil { log.Fatal(err) }\n}\n'

for name,(title,state,expr,behavior) in COMPONENTS.items():
    path=slug(name)
    write('examples/'+path+'/main.go', program(name,state,expr))
    entries=[e for e in API if e['Package']=='.' and e['Name'] in (name,name+'Props')]
    if name=='Text': entries += [e for e in API if e['Name']=='Label' and e['Package']=='.']
    if name=='Button': entries += [e for e in API if e['Name']=='TextButton' and e['Package']=='.']
    props=next(e for e in entries if e['Name']==name+'Props')
    table='| 字段 | 类型 | 用途、默认与约束 |\n| --- | --- | --- |\n'
    for line in props['Code'].splitlines()[1:]:
        if line.strip().startswith('//') or line.strip()=='}': continue
        m=re.match(r'\s*([A-Z]\w*(?:,\s*[A-Z]\w*)*)\s+(.+)',line)
        if m:
            for f in m[1].replace(' ','').split(','): table+=f'| `{f}` | `{m[2].strip()}` | {FIELDS[f]} |\n'
    write('components/'+path+'.md',f'# {name}：{title}\n\n用 `{name}` 构建{title}。本页包含可运行示例、完整属性及行为限制。\n\n## 完整示例\n\n先完成[环境准备](/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。\n\n<<< ../examples/{path}/main.go\n\n## 参数与 API\n'+''.join(code(e['Code']) for e in entries if e['Kind']=='func')+'\n'+table+'\n## 行为与边界\n\n'+behavior+'\n\n所有组件共享[身份、样式与指针规则](/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/api/components)。[示例运行方式](/examples/)包含构建验证方法。\n')

groups={'runtime':('应用生命周期','/api/application'), 'components':('组件关联类型','/components/'), 'style':('布局、样式与主题类型','/api/values'), 'tokens':('颜色与尺寸 Token','/api/theme'), 'resources':('字体、图片与图标资源','/api/resources-guide')}
def group(e):
    if e['File'] in ('app.go','dxui.go','window.go','theme_runtime.go','assign.go'): return 'runtime'
    if e['File'] in ('color_tokens.go','theme.go'): return 'tokens'
    if e['File'] in ('style.go',): return 'style'
    if e['File'] in ('font.go','icon_raster.go') or e['Package']=='internal/icondata': return 'resources'
    return 'components'

index='# 公开 API 索引\n\n以仓库快照为准；pre-v1 API 可能不兼容调整。每项链接到精确声明、源码契约与对应教程。包含根包、官方 icon、icon/catalog，以及 IconData 别名可调用的方法；不要求应用导入 internal。\n\n[更新日志](/api/coverage) · [图标全部名称](/api/icons)\n\n'
for key,(title,example) in groups.items():
    text=f'# {title}：完整声明\n\n[用途与示例]({example}) · [全部 API 索引](/api/)\n\n以下声明由 Go AST 提取，保留源码英文契约和字段注释，隐藏私有字段；`struct {{}}` 表示外部不可直接配置的内部表示，不表示可以用零值替代构造函数。常量块保留 iota 上下文。\n\n'
    seen=set()
    for e in API:
        if e['Package'] not in ('.','internal/icondata') or group(e)!=key: continue
        # Only alias-reachable internal API, no implementation helpers.
        if e['Package']=='internal/icondata' and e['Name'] not in ('Point','Rect','PathCommand','Data','Data.IsPacked','Data.Identity','Data.Paths','Path','PaintMode','PaintStroke','PaintFill'): continue
        name=e['Name'].replace('Data.', 'IconData.') if e['Package']=='internal/icondata' else e['Name']
        example_link = '/api/windows' if e['File']=='window.go' or name in ('WindowOptions', 'App.Invalidate') else example
        anchor='api-'+re.sub('[^a-z0-9-]','-',name.lower())
        index+=f'- [`{name}`](/api/{key}#{anchor}) — {title}；[示例]({example_link})\n'
        text+=f'## {name} {{#{anchor}}}\n\n'
        if e['Doc']: text+='```text\n'+e['Doc'].strip()+'\n```\n\n'
        if e['Code'] and e['Code'] not in seen:
            text+=code(e['Code']); seen.add(e['Code'])
        else: text+='类型和值见本页同组常量块。\n'
        text+=f'\n[对应示例]({example_link}) · 源文件 `{e["File"]}:{e["Line"]}`\n\n'
    write('api/'+key+'.md',text)
write('api/index.md',index)

icons=[e for e in API if e['Package']=='icon' and e['Kind']=='func']
names=re.findall(r'Name: "([^"]+)", GoName: "([^"]+)"', (Path(__import__('sys').argv[1])/'icon/catalog/catalog_generated.go').read_text())
iconpage='# 官方图标与完整目录\n\n`github.com/dxui-org/dxui/icon` 的每个构造器都无参数，返回不可变 `dxui.IconData`。用 `dxui.Icon(dxui.IconProps{Data: icon.Search()})` 显示；Size/Color/StrokeWidth 在视图上配置。[完整示例](/components/icon)。下表列出全部名称和别名；每行 Go 名称都可用作 `icon.名称()`。\n\n只有确实需要图标搜索器时才导入 `github.com/dxui-org/dxui/icon/catalog`：`catalog.Icons` 使整个图标库可达，增大产物。普通界面按需调用 icon 构造器。\n\n'
for e in API:
    if e['Package']=='icon/catalog': iconpage+=code(e['Code'])+'\n'+e['Doc']+'\n'
iconpage+='\n## 版本与数量常量\n\n`icon.Version` 是上游 Lucide 版本字符串；`icon.Count` 包含规范名称和别名；`icon.CanonicalCount` 仅统计规范名称。这些是编译期元数据，没有参数，也不触发图标加载。可用 `fmt.Println(icon.Version, icon.Count, icon.CanonicalCount)` 显示当前编译基线。\n'
for e in API:
    if e['Package']=='icon' and e['Kind']!='func': iconpage+=code(e['Code'])
iconpage+='\n## 全部构造器\n\n'
iconpage+='```go\nfor _, entry := range catalog.Icons {\n    // Name 为 kebab-case，GoName 为精确 Go 构造器名。\n    if entry.Name == "search" {\n        view := dxui.Icon(dxui.IconProps{Data: entry.Icon()})\n        _ = view\n    }\n}\n```\n\n| 名称 | Go 构造器（返回 dxui.IconData） |\n| --- | --- |\n'
iconpage+=''.join(f'| `{n}` | `icon.{g}()` |\n' for n,g in names)
write('api/icons.md',iconpage)
write('components/index.md','# 组件参考\n\n每页包含可独立运行的完整程序、全部 Props 字段、交互行为和限制。`Label` 合并在 Text 页，`TextButton` 合并在 Button 页。\n\n'+''.join(f'- [{n}：{v[0]}](/components/{slug(n)})\n' for n,v in COMPONENTS.items()))
write('examples/index.md','# 运行完整示例\n\n每个组件页引用本站 `docs/examples/<组件>/main.go` 的同一份代码，避免展示代码与编译代码分叉。界面 API 仅使用根包与官方 icon 资源；Image/Avatar 示例自行创建像素，不需要外部文件。\n\n在任意空目录执行快速开始的 `go mod init`、`go get`，把所选示例复制为 `main.go` 后 `go run .`。这些是原生窗口程序，不是在浏览器中运行的 Go。\n\n- [多窗口与跨窗口状态更新](/guide/multi-window#完整示例)\n- [主窗口与子窗口控制](/api/windows#完整操作示例)\n\n'+''.join(f'- [{n} 完整示例](/components/{slug(n)}#完整示例)\n' for n in COMPONENTS)+'\n维护者在网站根目录执行 `python scripts/check-docs.py <dxui源码路径>`，会逐个编译完整示例并核对公开 API 快照。编译成功不能替代本机窗口交互验收。\n')

# Keep the displayed snippets identical to the compiler inputs.
import subprocess
subprocess.run(['go', 'run', './scripts/format-examples.go', 'docs'], cwd=ROOT, check=True)
