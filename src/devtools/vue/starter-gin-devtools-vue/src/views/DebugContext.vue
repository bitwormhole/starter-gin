<template>
  <div class="about">
    <h1>环境信息</h1>

    <el-button
      class="btn-refresh el-icon-refresh"
      type="primary"
      @click="refresh"
    >
      刷新
    </el-button>

    <el-collapse v-model="activeName" accordion>
      <!-- arguments  -->

      <el-collapse-item title="Arguments" name="1">
        <el-table :data="myArguments">
          <el-table-column
            prop="index"
            label="Index"
            width="100px"
          ></el-table-column>
          <el-table-column prop="value" label="Value"></el-table-column>
        </el-table>
      </el-collapse-item>

      <!-- env  -->
      <el-collapse-item title="Environment" name="2">
        <el-table :data="myEnvironment">
          <el-table-column
            prop="index"
            label="Index"
            width="100px"
          ></el-table-column>
          <el-table-column prop="name" label="Name"></el-table-column>
          <el-table-column prop="value" label="Value"></el-table-column>
        </el-table>
      </el-collapse-item>

      <!-- properties  -->
      <el-collapse-item title="Properties" name="3">
        <el-table :data="myProperties">
          <el-table-column
            prop="index"
            label="Index"
            width="100px"
          ></el-table-column>
          <el-table-column prop="name" label="Name"></el-table-column>
          <el-table-column prop="value" label="Value"></el-table-column>
        </el-table>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
export default {
  name: "debug-context",

  data() {
    return {
      myArguments: [],
      myEnvironment: [],
      myProperties: [],
      activeName: 1,
    };
  },

  mounted() {
    this.refresh();
  },

  methods: {
    refresh() {
      this.$store
        .dispatch("api/httpGet", { type: "context" })
        .then((res) => {
          this.handleHttpOk(res.data);
        })
        .catch((err) => {});
    },

    handleHttpOk(data) {
      this.myArguments = this.convertStringArrayToList(data.arguments);
      this.myEnvironment = this.convertTableToList(data.environment);
      this.myProperties = this.convertTableToList(data.properties);
    },

    convertTableToList(table) {
      let list = [];
      let keys = [];
      for (var key in table) {
        keys.push(key);
      }
      keys.sort();
      for (var i in keys) {
        let key = keys[i];
        let val = table[key];
        list.push({
          index: i,
          name: key,
          value: val,
        });
      }
      return list;
    },

    convertStringArrayToList(strings) {
      let list = [];
      for (var i in strings) {
        let key = i + "";
        let val = strings[i];
        list.push({ key, index: key, value: val });
      }
      return list;
    },
  },
};
</script>

<style scoped>
.btn-refresh {
  margin: 30px;
}
</style>