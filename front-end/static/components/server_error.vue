<script>
import fire from "./event_bus";

export default {
  data() {
    return {
      value: ""
    }
  },
  created() {
    fire.showServerError = this.showServerError
  },
  methods: {
    showServerError(msg) {
      this.value = msg
    },
    hide() {
      this.value = ""
    }
  }
}
</script>

<template>
  <div>
    <div v-if="value.length > 0">
      <button @click="hide()">&#x2713;</button>
      &nbsp;
      <strong>Error:</strong>&nbsp;{{ value }}
    </div>
  </div>
</template>

<style scoped>

</style>