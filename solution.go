package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
import (
	"math"
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:

// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type intCustomType int

const (
	SidesTriangle = 3
	SidesSquare = 4
	SidesCircle = 0
  )
func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	
	if sidesNum == intCustomType(SidesTriangle){
		return sideLen * math.Sqrt(3) * sideLen / float64(4)
	}else if sidesNum == intCustomType(SidesSquare){
		return sideLen*sideLen
	}else if sidesNum == intCustomType(SidesCircle){
		return math.Pi*sideLen * sideLen
	}else{
		return 0
	}
}

