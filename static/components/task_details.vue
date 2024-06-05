<script>
import * as api from "./api";
import fire from "./event_bus";

const NO_TASK = {"t_id": -1, "p_id": -1, "t_date": null, "t_subject": null, "t_priority": -1, "t_comments": null}

export default {
  data() {
    return {
      current_subject: null,
      current_task: NO_TASK,
      task_error: null,
      show_task_details: false,
    }
  },
  created() {
    fire.renderTaskDetails = this.renderTaskDetails
    fire.hideTaskDetails = this.hideTaskDetails
  },
  methods: {
    hideTaskDetails() {
      this.show_task_details = false
    },
    renderTaskDetails(t_id) {
      fetch("/api/tasks/" + t_id)
          .then(async (resp) => {
            if (resp.status === 200) {
              let task = await resp.json()
              this.current_subject = task.t_subject;
              this.current_task = task;
              this.task_error = null;
              this.show_task_details = true
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    taskUpdate() {
      if (!isNaN(this.current_task.t_priority)) {
        this.current_task.t_priority = parseInt(this.current_task.t_priority);
      }
      let json = JSON.stringify(this.current_task)
      let p_id = this.current_task.p_id
      let t_id = this.current_task.t_id
      fetch("/api/tasks/" + t_id, {
        method: 'put',
        headers: api.JSON_HEADERS,
        body: json
      })
          .then(async (resp) => {
            if (resp.status === 200) {
              fire.renderProjectTasks(p_id);
              this.renderTaskDetails(t_id);
            } else {
              let text = await resp.text()
              this.task_error = (resp.status + "\n" + text);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    taskDelete() {
      let p_id = this.current_task.p_id
      let t_id = this.current_task.t_id
      fetch("/api/tasks/" + t_id, {
        method: "delete"
      })
          .then(async (resp) => {
            if (resp.status === 204) {
              this.hideTaskDetails()
              fire.renderProjects(); // update tasks count
              fire.renderProjectTasks(p_id);
            } else {
              let text = await resp.text()
              alert(resp.status + "\n" + text);
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
  <div v-if="show_task_details">
    <div class="title" id="subj">
      {{ current_subject }}
    </div>
    <table>
      <tr>
        <td class="form-label">Date</td>
        <td>
          <label>
            <input type="text" v-model="current_task.t_date"/>
          </label>
        </td>
      </tr>
      <tr>
        <td class="form-label">Subject</td>
        <td class="w100">
          <label>
            <input type="text" v-model="current_task.t_subject"/>
          </label>
        </td>
      </tr>
      <tr>
        <td class="form-label">Priority</td>
        <td>
          <label>
            <input type="text" v-model="current_task.t_priority"/>
          </label>
        </td>
      </tr>
      <tr>
        <td colspan="2">
          <label>
            <textarea cols="40" rows="10" v-model="current_task.t_comments"></textarea>
          </label>
        </td>
      </tr>
      <tr>
        <td>
          <table class="controls">
            <tr>
              <td class="w1">
                <a href="#" @click="taskUpdate()">
                  <input type="button" value="✓"/>
                </a>
              </td>
              <td class="w1">
                <a href="#" @click="taskDelete()">
                  <input type="button" value="x"/>
                </a>
              </td>
              <td class="w1">
                <a href="#" @click="hideTaskDetails()">
                  <input type="button" value="&lt;"/>
                </a>
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
    <div style="color: crimson;max-width: 310px;">
      {{ task_error }}
    </div>
  </div>
</template>

<style scoped>

</style>