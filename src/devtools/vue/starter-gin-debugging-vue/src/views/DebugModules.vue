<template>
  <div class="about">
    <h1>模块信息</h1>

    <el-button
      class="btn-refresh el-icon-refresh"
      type="primary"
      @click="refresh"
    >
      刷新
    </el-button>

    <el-table :data="myModules">
      <el-table-column prop="Index" label="序号"></el-table-column>
      <el-table-column prop="Name" label="名称"></el-table-column>
      <el-table-column prop="Version" label="版本"></el-table-column>
      <el-table-column prop="Revision" label="版次"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "debug-modules",

  data() {
    return {
      myModules: [],
    };
  },

  mounted() {
    this.refresh();
  },

  methods: {
    refresh() {
      this.$store
        .dispatch("api/httpGet", { type: "modules" })
        .then((res) => {
          this.handleHttpOk(res.data);
        })
        .catch((err) => {});
    },

    handleHttpOk(data) {
      let list = data.modules;
      for (var i in list) {
        let item = list[i];
        item.key = item.Name + "@" + item.Version;
      }
      list.sort((a, b) => {
        return this.cmpstr(a.Index, b.Index);
      });
      this.myModules = list;
    },

    cmpstr(a, b) {
      let n1 = Number(a);
      let n2 = Number(b);
      return n1 - n2;
    },
  },
};
</script>

<style scoped>
.btn-refresh {
  margin: 30px;
}
</style>