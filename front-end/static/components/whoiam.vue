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
        this.value = text + ", npm, vue " + Vue.version
        document.title = 'SDM Todo-App, npm, vue ' + Vue.version
      })
    },
  }
}
</script>

<template>
  <table>
    <tr>
      <td>
        <h2>
          {{ value }}
        </h2>
      </td>
      <td v-if="is_sqlx">
        <a href="swagger/index.html">Swagger</a>
      </td>
    </tr>
  </table>
</template>

<style>

</style>
