<script>
import my_event_bus from "./my_event_bus";

const PROJECT = {"p_id": -1, "p_name": null, "p_tasks_count": -1}

import * as shared from './shared.js'

export default {
  data() {
    return {
      projects: [PROJECT],
      p_name: null,
    }
  },
  created() {
    my_event_bus.renderProjects = this.renderProjects;
  },
  methods: {
    renderProjects() {
      fetch("/api/projects")
          .then(async (resp) => {
            if (resp.status === 200) {
              this.projects = await resp.json()
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    renderProjectDetails(p_id) {
      my_event_bus.renderProjectDetails(p_id)
    },
    projectCreate() {
      let json = JSON.stringify({"p_name": this.p_name})
      fetch("/api/projects", {
        method: 'post',
        headers: shared.JSON_HEADERS,
        body: json
      })
          .then(async (resp) => {
            if (resp.status === 201) {
              this.renderProjects();
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
  }
}
</script>

<template>
  <table>
    <tr>
      <td>
        <table class="section">
          <tr>
            <th>Project</th>
            <th>Tasks Count</th>
          </tr>
          <tr v-for="project in projects">
            <td>
              <a href="#" @click="renderProjectDetails(project.p_id)">
                {{ project.p_name }}
              </a>
            </td>
            <td style="text-align: center;width: 80px;">
              {{ project.p_tasks_count }}
            </td>
          </tr>
        </table>
      </td>
    </tr>
    <tr>
      <td>
        <table class="controls">
          <tr>
            <td>
              <label>
                <input type="text" v-model="p_name"/>
              </label>
            </td>
            <td class="w1">
              <a href="#" @click="projectCreate()">
                <input type="button" value="+"/>
              </a>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</template>

<style scoped>

</style>