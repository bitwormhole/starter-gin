<template>
  <div class="about">
    <h1>Starter Gin 开发助手</h1>

    <mods />
  </div>
</template>

<script>
import mods from "./DebugModules.vue";

export default {
  name: "about",
  components: { mods },
};
</script>