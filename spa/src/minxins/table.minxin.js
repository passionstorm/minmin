const t = {
  data(){
    return {
      data: [],
      pageParams: {
        pageNum: 1,
        pageSize: 100,
        page: true,
      },
      filterParams: {
        id: '',
      },
      tableList: [],
      filterList: [],
      editData: {},
    }
  }
}