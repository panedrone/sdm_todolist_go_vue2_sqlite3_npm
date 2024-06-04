<template>
  <table>
    <tr>
      <td>
        <h2>
          SDM, {{ value }}
        </h2>
      </td>
      <td v-if="is_sqlx">
        <a href="/swagger/index.html">Swagger</a>
      </td>
    </tr>
  </table>
</template>

<script>
import Vue from "vue";
import my_event_bus from "./my_event_bus";

export default {
  data() {
    return {
      value: null,
      is_sqlx: false
    }
  },
  created() {
    my_event_bus.renderWhoIAm = this.renderWhoIAm
  },
  methods: {
    renderWhoIAm() {
      fetch("/api/whoiam")
          .then(async (resp) => {
            if (resp.status === 200) {
              this.value = await resp.text()
              this.value += ", npm, vue " + Vue.version
              document.title = 'Todo-App, Go, NPM, Vue ' + Vue.version
              if (this.value.includes('sqlx')) {
                this.is_sqlx = true
              }
            } else {
              let j = await resp.text()
              console.log(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
  },
  // mounted() {
  //   this.askWhoIAm()
  // }
}
</script>

<style>

</style>