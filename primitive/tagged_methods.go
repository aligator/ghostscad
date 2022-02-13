// Do not edit. Automatically generated on Sun Feb 13 18:56:37 CET 2022.

package primitive

import ()

func (o *Cube) SetCenter(val bool) *Cube {
	o.Center = val
	return o
}

func (o *Cylinder) SetRBottom(val float64) *Cylinder {
	o.RBottom = val
	return o
}

func (o *Cylinder) SetCenter(val bool) *Cylinder {
	o.Center = val
	return o
}

func (o *Import) SetConvexity(val int) *Import {
	o.Convexity = val
	return o
}

func (o *Import) SetLayer(val string) *Import {
	o.Layer = val
	return o
}

func (o *LinearExtrusion) SetCenter(val bool) *LinearExtrusion {
	o.Center = val
	return o
}

func (o *LinearExtrusion) SetConvexity(val uint16) *LinearExtrusion {
	o.Convexity = val
	return o
}

func (o *LinearExtrusion) SetTwist(val uint16) *LinearExtrusion {
	o.Twist = val
	return o
}

func (o *LinearExtrusion) SetSlices(val uint16) *LinearExtrusion {
	o.Slices = val
	return o
}

func (o *LinearExtrusion) SetScale(val float64) *LinearExtrusion {
	o.Scale = val
	return o
}

func (o *LinearExtrusion) SetFn(val uint16) *LinearExtrusion {
	o.Fn = val
	return o
}

func (o *Polygon) SetPaths(val [][]int) *Polygon {
	o.Paths = val
	return o
}

func (o *Polygon) SetConvexity(val int) *Polygon {
	o.Convexity = val
	return o
}

func (o *Polyhedron) SetConvexity(val int) *Polyhedron {
	o.Convexity = val
	return o
}

func (o *RotationExtrusion) SetAngle(val float64) *RotationExtrusion {
	o.Angle = val
	return o
}

func (o *RotationExtrusion) SetConvexity(val uint16) *RotationExtrusion {
	o.Convexity = val
	return o
}

func (o *Surface) SetCenter(val bool) *Surface {
	o.Center = val
	return o
}

func (o *Surface) SetInvert(val bool) *Surface {
	o.Invert = val
	return o
}

func (o *Surface) SetConvexity(val int) *Surface {
	o.Convexity = val
	return o
}

func (o *Cylinder) SetFa(val float64) *Cylinder {
	o.Circular.SetFa(val)
	return o
}

func (o *Cylinder) SetFs(val float64) *Cylinder {
	o.Circular.SetFs(val)
	return o
}

func (o *Cylinder) SetFn(val uint16) *Cylinder {
	o.Circular.SetFn(val)
	return o
}

func (o *LinearExtrusion) Add(items ...Primitive) *LinearExtrusion {
	o.Items.Add(items...)
	return o
}

func (o *RotationExtrusion) SetFa(val float64) *RotationExtrusion {
	o.Circular.SetFa(val)
	return o
}

func (o *RotationExtrusion) SetFs(val float64) *RotationExtrusion {
	o.Circular.SetFs(val)
	return o
}

func (o *RotationExtrusion) SetFn(val uint16) *RotationExtrusion {
	o.Circular.SetFn(val)
	return o
}

func (o *RotationExtrusion) Add(items ...Primitive) *RotationExtrusion {
	o.Items.Add(items...)
	return o
}

func (o *Scale) Add(items ...Primitive) *Scale {
	o.Items.Add(items...)
	return o
}

func (o *Sphere) SetFa(val float64) *Sphere {
	o.Circular.SetFa(val)
	return o
}

func (o *Sphere) SetFs(val float64) *Sphere {
	o.Circular.SetFs(val)
	return o
}

func (o *Sphere) SetFn(val uint16) *Sphere {
	o.Circular.SetFn(val)
	return o
}

func (o *Translation) Add(items ...Primitive) *Translation {
	o.Items.Add(items...)
	return o
}

func (o *Cube) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Cube) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Cube) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Cube) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Cube) Prefix() string {
	return o.prefix
}

func (o *Cylinder) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Cylinder) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Cylinder) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Cylinder) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Cylinder) Prefix() string {
	return o.prefix
}

func (o *Import) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Import) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Import) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Import) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Import) Prefix() string {
	return o.prefix
}

func (o *LinearExtrusion) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *LinearExtrusion) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *LinearExtrusion) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *LinearExtrusion) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *LinearExtrusion) Prefix() string {
	return o.prefix
}

func (o *List) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *List) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *List) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *List) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *List) Prefix() string {
	return o.prefix
}

func (o *Polygon) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Polygon) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Polygon) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Polygon) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Polygon) Prefix() string {
	return o.prefix
}

func (o *Polyhedron) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Polyhedron) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Polyhedron) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Polyhedron) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Polyhedron) Prefix() string {
	return o.prefix
}

func (o *RotationExtrusion) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *RotationExtrusion) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *RotationExtrusion) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *RotationExtrusion) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *RotationExtrusion) Prefix() string {
	return o.prefix
}

func (o *Scale) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Scale) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Scale) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Scale) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Scale) Prefix() string {
	return o.prefix
}

func (o *Sphere) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Sphere) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Sphere) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Sphere) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Sphere) Prefix() string {
	return o.prefix
}

func (o *Surface) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Surface) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Surface) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Surface) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Surface) Prefix() string {
	return o.prefix
}

func (o *Translation) Disable() Primitive {
	o.prefix = "*"
	return o
}

func (o *Translation) ShowOnly() Primitive {
	o.prefix = "!"
	return o
}

func (o *Translation) Highlight() Primitive {
	o.prefix = "#"
	return o
}

func (o *Translation) Transparent() Primitive {
	o.prefix = "%"
	return o
}

func (o *Translation) Prefix() string {
	return o.prefix
}
