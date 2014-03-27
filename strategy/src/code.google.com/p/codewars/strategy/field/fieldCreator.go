package field


import (
	"math/rand"
	"time"
	"math"
)

var fieldSize = 50
var MaxPiek = 50
var MinPiek = 30


var gaus =[]int{1,4,7,10,7,4,1,
				4,16,26,36,26,16,4,
				7,26,41,60,41,26,7,
				10,36,60,89,60,36,10,
				7,26,41,60,41,26,7,
				4,16,26,36,26,16,4,
				1,4,7,10,7,4,1}

var gausSum = 0.0;
var gausFieldSz = 0;

func GetFieldSize() int{
	return fieldSize
}

func ChangeFieldSize(size int){
	fieldSize = size
	gausSize := fieldSize/10
	gausSize += (gausSize+1)%2//always odd
	gaus = make([]int,gausSize*gausSize)
	mdl:=gausSize/2
	maxD := math.Sqrt(float64(mdl*mdl+mdl*mdl))
	
	idx := 0
	for x:=0;x<gausSize;x++{
		for y:=0;y<gausSize;y++{
			curD:=math.Sqrt(float64((mdl-x)*(mdl-x)+(mdl-y)*(mdl-y)))
			gaus[idx]=90-int(curD/maxD*90.0)
			idx++
		}
	}
	initGaus()
}

func initGaus(){
	curSum := 0
	
	
	for i:=0;i<len(gaus);i++ {
		curSum+=gaus[i]
	}
	gausSum = float64(curSum)
	gausFieldSz = int(math.Sqrt(float64(len(gaus)))/2.0)
}

func init() {
	initGaus()
}

func CreateField() ([][]int) {

	rand.Seed(int64(time.Now().Nanosecond()))//Seed != 1 setzen
	myField := make([][]int, fieldSize)
	for y:= range myField{
		myField[y] = make([]int, fieldSize)
	}
	pieks := MinPiek+rand.Intn(MaxPiek-MinPiek)//maximal MaxPiek berge
	for idx:=0;idx<pieks;idx++{
		x := rand.Intn(fieldSize)
		y := rand.Intn(fieldSize)
		createIsl(myField,y,x,50 + rand.Intn(250))
	}
	
	linearize(myField)
	
	
	return myField;
}

func createIsl(field [][]int,y int, x int, height int){
	field[y][x] = height
	islSz := fieldSize /50
	islSz += (islSz+1)%2//always odd
	islSz /=2
	stX := x-islSz
	if stX < 0{
		stX = 0
	}
	stY := y-islSz
	if stY < 0{
		stY = 0
	}
	for ;stY<y+islSz;stY++{
		if stY >= len(field){
			break
		}
		for ;stX<x+islSz;stX++{
			if stX >= len(field[stY]){
				break
			}
			field[stY][stX] = height
		}
	}
}

func linearize(field [][]int){
	for idx:=0;idx<1;idx++{//5 runden glätten
		linearizeSingle(field)
	}
}

func linearizeSingle(field [][]int){
	for y:=0;y<len(field);y++{
		for x:=0;x<len(field[y]);x++{
			middl := getMiddle(field,x,y)
			if middl > field[y][x]{
				field[y][x] = middl //nur setzen, wenn größer
			}
			field[y][x] = middl
		}
	}
}

func getMiddle(field [][]int, x int, y int) (int){
	gausIdx:=0
	m := 0.0
	stY := y-gausFieldSz
	mxY := y+gausFieldSz
	for curY:=stY;curY<=mxY;curY++{
		stX := x-gausFieldSz
		mxX := x+gausFieldSz
		for curX:=stX;curX<=mxX;curX++{
			curGaus := float64(gaus[gausIdx]);
			gausIdx++
			if curY < 0{
				continue;
			}
			if curY > len(field)-1{
				continue;
			}
			if curX < 0{
				continue;
			}
			if curX > len(field[curY])-1{
				continue;
			}
			curVal := float64(field[curY][curX])
			m=m+(curVal*curGaus/gausSum)
			//if curX != x || curY != y{
				//curVal := field[curY][curX]
				//if curVal == 0{
				//	field[curY][curX] = 1
				//	curVal = 1
				//}
				//m = m+(float64(field[curY][curX]))
			//}
		}
	}
	//debug output 
	//fmt.Printf("m=%g;c=%d",m,gausIdx)
	return int(m)
}
