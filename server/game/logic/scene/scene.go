package scene

import (
	"server/conf/game"
	"strings"
)

var sceneObj *Scene

type Scene struct {
	sceneMap map[string]ScenePvp
}

func NewScene() *Scene {
	if sceneObj == nil {
		sceneObj = &Scene{make(map[string]ScenePvp)}
	}
	return sceneObj
}
func (this *Scene) GetAllScene() map[string]ScenePvp {
	return this.sceneMap
}
func (this *Scene) GetSceneByKey(key string) ScenePvp {
	scene, ok := this.sceneMap[key]
	if !ok {
		return nil
	}
	return scene
}
func (this *Scene) SetScene(key string, Scene ScenePvp) {
	this.sceneMap[key] = Scene
}
func (this *Scene) GetOrCreateScenePvp(key string) ScenePvp {
	SceneObj := this.GetSceneByKey(key)
	if SceneObj == nil {
		SceneObj = this.CreateScene(key)
	}
	return SceneObj
}
func (this *Scene) CreateScene(mapId string) ScenePvp {
	missConfs := game.GetConf("mission")
	missConf := missConfs[mapId].(map[string]interface{})
	//monsterConf := game.GetConf("monster")
	sceneObj := NewScenePvpNormal()
	sceneObj.SceneId = "main_" + mapId
	this.CreateBoss(missConf, sceneObj)
	this.CreateMonster(missConf, sceneObj)
	this.SetScene(mapId, sceneObj)
	return sceneObj
}
func (this *Scene) CreateBoss(missConf map[string]interface{}, sceneObj *ScenePvpNormal) {
	bossIds := strings.Split(missConf["bossId"].(string), ",")
	for _, bossId := range bossIds {
		bossRole := NewPvpRole().CreateMonster(bossId, PVP_ROLE_BOSS)
		sceneObj.AddMonster(bossRole)
	}
}
func (this *Scene) CreateMonster(missConf map[string]interface{}, sceneObj *ScenePvpNormal) {
	bossIds := strings.Split(missConf["monsterId"].(string), ",")
	for _, bossId := range bossIds {
		bossRole := NewPvpRole().CreateMonster(bossId, PVP_ROLE_MONSTER)
		sceneObj.AddMonster(bossRole)
	}
}
