一般后面可接受三个[[../../../_resources/documents/CSS/长度或百分比|尺寸]]大小，分别表示 x，y 偏移长度和阴影半径（模糊度） `blur-radius`，以及一个[[../../../_resources/documents/CSS/颜色|颜色]]。

同一个阴影之间各参数使用空格分割，同时添加的多个阴影可用 `,` 分割。存在多个阴影时，按 z 轴向下叠加。

对于 `box-shadow`：
- 可以额外接受一个[[../../../_resources/documents/CSS/长度或百分比|尺寸]]大小表示 `spread-radius`，为正时阴影扩大，否则阴影缩小。
- 可以额外接受一个 `inset` 关键字，如果存在则阴影在盒子之内，盒子实际大小会有所缩小