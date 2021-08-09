<template>
  <div class="about">
    <h1>HTTP 请求统计</h1>

    <el-button class="btn-reset el-icon-reset" type="default" @click="reset">
      复位
    </el-button>

    <el-button
      class="btn-refresh el-icon-refresh"
      type="primary"
      @click="refresh"
    >
      刷新
    </el-button>

    <span>
      自动刷新 : <el-switch v-model="enableRefreshTimer"> </el-switch> {{
        timerCount
      }}
      次
    </span>

    <el-table :data="myEndpoints">
      <el-table-column prop="index" label="序号"></el-table-column>
      <el-table-column prop="Method" label="方法"></el-table-column>
      <el-table-column prop="Path" label="路径"></el-table-column>
      <el-table-column prop="Count" label="请求计数"></el-table-column>
      <el-table-column
        prop="TimeSpan"
        label="持续时间"
        :formatter="formatTimeSpan"
      ></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "debug-requests",

  data() {
    return {
      myEndpoints: [],
      timerCount: 0,
      timerInterval: 1000,
      timerThreadId: 0,
      enableRefreshTimer: false,
      netWorking: false,
    };
  },

  watch: {
    enableRefreshTimer(enable) {
      if (enable) {
        this.startTimer();
      } else {
        this.stopTimer();
      }
    },
  },

  mounted() {
    this.refresh();
  },

  deactivated() {
    this.stopTimer();
  },

  methods: {
    refresh() {
      this.netWorking = true;
      this.$store
        .dispatch("api/httpGet", { type: "requests" })
        .then((res) => {
          this.handleHttpOk(res.data);
          this.netWorking = false;
        })
        .catch((err) => {
          this.netWorking = false;
        });
    },

    handleHttpOk(data) {
      let src = data.requests;
      for (var i in src) {
        let item = src[i];
        item.index = i;
        item.key = item.Method + ":" + item.Path;
        item.TimeSpan = item.TimeEnd - item.TimeBegin;
      }
      this.myEndpoints = src;
    },

    reset() {
      this.$store
        .dispatch("api/httpDelete", { type: "requests" })
        .then((res) => {})
        .catch((err) => {});
    },

    formatTimeSpan(row, column, cellValue, index) {
      let sec = cellValue / 1000;
      return sec + "秒";
    },

    onTimer(tid) {
      if (tid != this.timerThreadId) {
        return;
      }
      let self = this;
      let interval = this.timerInterval;
      // for next time
      setTimeout(() => {
        self.onTimer(tid);
      }, interval);

      // do refresh
      if (this.netWorking) {
        return;
      }
      this.refresh();
      this.timerCount++;
    },

    startTimer() {
      let now = new Date();
      let tid = now.getTime();
      this.timerThreadId = tid;
      this.timerCount = 0;
      this.onTimer(tid);
    },

    stopTimer() {
      this.timerThreadId = 0;
    },
  },
};
</script>

<style scoped>
.btn-refresh {
  margin: 30px;
}
</style>