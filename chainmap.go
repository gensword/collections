package collections

type ChainMap struct {
	Map  map[string]interface{}
	Maps []map[string]interface{}
}

func NewChainMap(maps ...map[string]interface{}) (chainMap *ChainMap) {
	chainMap = new(ChainMap)
	chainMap.Map = make(map[string]interface{})
	chainMap.Maps = make([]map[string]interface{}, 0)
	for _, item := range maps {
		for k, v := range item {
			if _, ok := chainMap.Map[k]; !ok {
				chainMap.Map[k] = v
			}
		}
		chainMap.Maps = append(chainMap.Maps, item)
	}
	return
}

func (cm *ChainMap) NewChild(child map[string]interface{}) {
	for k, v := range child {
		if _, ok := cm.Map[k]; !ok {
			cm.Map[k] = v
		}
	}
	cm.Maps = append(cm.Maps, child)
}

func (cm *ChainMap) Parents() *ChainMap {
	parentsChainMap := NewChainMap(cm.Maps[1:len(cm.Maps)]...)
	return parentsChainMap
}

func (cm *ChainMap) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range cm.Map {
		keys = append(keys, k)
	}
	return keys
}

func (cm *ChainMap) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, v := range cm.Map {
		values = append(values, v)
	}
	return values
}
