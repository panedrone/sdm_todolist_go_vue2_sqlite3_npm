<script>
import fire from "./event_bus";
import * as api from './api'
import * as api from "./api";

const PROJECT = {"p_id": -1, "p_name": null, "p_tasks_count": -1}

export default {
  data() {
    return {
      projects: [PROJECT],
      p_name: null,
    }
  },
  created() {
    fire.renderProjects = this.renderProjects;
  },
  methods: {
    renderProjects() {
      api.fetchJsonArray("api/projects", (arr) => {
        this.projects = arr
      })
    },
    renderProjectDetails(p_id) {
      fire.renderProjectDetails(p_id)
    },
    projectCreate() {
      let json = JSON.stringify({"p_name": this.p_name})
      api.postJson201("api/projects", json, () => {
        this.renderProjects();
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