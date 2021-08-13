<template>
  <div class="about">
    <h1>组件信息</h1>

    <el-button
      class="btn-refresh el-icon-refresh"
      type="primary"
      @click="refresh"
    >
      刷新
    </el-button>

    <el-table :data="items">
      <el-table-column prop="index" label="序号"></el-table-column>
      <el-table-column prop="ID" label="ID"></el-table-column>
      <el-table-column prop="Class" label="Class"></el-table-column>
      <el-table-column prop="Aliases" label="Aliases"></el-table-column>
      <el-table-column prop="Scope" label="Scope"></el-table-column>
      <el-table-column prop="Type" label="Type"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "debug-components",

  data() {
    return {
      items: [],
    };
  },

  mounted() {
    this.refresh();
  },

  methods: {
    refresh() {
      this.$store
        .dispatch("api/httpGet", { type: "components" })
        .then((res) => {
          this.handleHttpOk(res.data);
        })
        .catch((err) => {});
    },

    handleHttpOk(data) {
      let list = data.components;
      for (var i in list) {
        let item = list[i];
        item.index = i;
      }
      list.sort((a, b) => {
        return this.cmpstr(a.ID, b.ID);
      });
      this.items = list;
    },

    cmpstr(a, b) {
      let n1 = "" + a;
      let n2 = "" + b;
      return n1 > n2;
    },
  },
};
</script>

<style scoped>
.btn-refresh {
  margin: 30px;
}
</style>