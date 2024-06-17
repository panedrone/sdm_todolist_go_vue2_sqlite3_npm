<script>
import Vue from "vue";

import fire from "./event_bus";
import {getText} from "./api";

export default {
  data() {
    return {
      value: "",
      is_sqlx: false
    }
  },
  created() {
    fire.renderWhoIAm = this.renderWhoIAm
  },
  methods: {
    renderWhoIAm() {
      getText('api/whoiam', (text) => {
        if (!text) {
          return
        }
        if (text.includes('sqlx')) {
          this.is_sqlx = true
        }
        this.value = text + ", npm, babel, vue " + Vue.version
        document.title = 'SDM Todo, npm, babel, vue ' + Vue.version
      })
    },
  }
}
</script>

<template>
  <div>
    {{ value }}<span v-if="is_sqlx">, <a href="swagger/index.html">swagger</a></span>
  </div>
</template>

<style>

</style>
