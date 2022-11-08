
# Quinne 
Quinne is a framework to create a new Domain-Specific programming language(DSL)

***Goal***:  To make it possible and easier to create a new domain specific language in a short time, with this newly created language in hand, development teams can have these advantages:
1. ***Efficient comunication***:  Developers can communicate with other teams in this new, domain-specific, easier to understand language, most misunderstandings can be eliminated. 
2. ***Clean task decomposition***: With DSL, clean logic branches, less-parameter function calls, it's natural to express complicated bussiness logic by directional-programming. Abstractions, task decompositions, development time estimations can be done at this stage.
3. ***Formalized protocal***: Protocals(Rules) can be agreed when logic has been discribed in DSL, any further changes can be applied to this new language, thus easier to know if any substantial code changes are required or exactly who to make these changes, developers and non-developers can reach agreement faster.
4.  ***Less documentations and tests***: With DSL itself expressing bussioness logic already, less documentation is required, overlap test cases can be significantly eliminated.  
5. ***Customized Grammar***: With Golang's basic grammar combined with Python style try-catch, lamba grammar, developer's can write shorter programs with enhanced functionalities by using functional programming.

**More Readings**:
- [Looking at programming language from a different aspect](https://medium.com/@quinnkunfang_5420/looking-at-programming-language-from-a-different-aspect-9e4544047c1e)


## Getting started

### Getting Quinne
```bash
go get -u github.com/Quinn-Fang/quinne
```
## Examples
### 1. Graphs:
#### Electronic circuit:
```go
package samples_001

func main() {
	create_battery()
	create_switch()
	if SWITCH_ON {
		create_bulb("bulb-1")
	} else {
		create_bulb("bulb-2")
	}
}
```

#### Trees(Database index storage)
```go
package graph

func main() {
	searchForTreeNodeWithValue(36)
	if treeNodeExists {
		if rearrangeRequired {
			rearrangeTreeStructure()
			insertNewValue()
		} else {
			insertNewValue()
		}
	} else {
		createTreeNode()
	}
}

```

### 2. Directional programming
#### Route planner
```go
package directional

func main() {
	activateVehicle("Mercedes_1")
	weather := GetWeatherStatus()
	if weather == "raining" {
		showNotificationOnDevice("IPhone_device_1", "it's raining heavily, bring umbrella")
		driveToDestination("pandas_1")
	} else if weather == "snowing" {
		showNotificationOnDevice("Apple_watch_device_2", "it's snowing, wear more clothes")
		driveToBackDoor()
		openBackDoor()
		driveToCompany()
	}
}

```

## Docs:
- [Introduction: ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/info.md)
- [Examples: ](https://github.com/Quinn-Fang/quinne/blob/master/Documents/info.md)

