export default class listUtils {
  static logger(str: string) {
    console.log(str)
  }

  static joinItemsAndActions(items: List<any>, actions: List<any>) {
    const resList: any[] = []
    for (item of items) {
      tempItem = {
        title: item.title,
        description: item.description,
        definition: {
          item: item.description.item,
          useEffects: {},
          useCriticalEffects: {},
          equipEffects: {}
        }
      }
      for (equipEffect of item.definition.equipEffects) {
        resList.push()
      }
    }
  }
}
