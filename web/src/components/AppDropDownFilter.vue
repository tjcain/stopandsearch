<template>
  <div class="box">
    <nav class="level">
      <div class="level-left">
        <div class="level-ite">
          <p class="menu-label">Age Range</p>
        </div>
      </div>
      <div class="level-right">
        <div class="level-item">
          <b-switch :value="true" type="is-success"></b-switch>
        </div>
        <div class="level-item">
          <a v-if="!dropped" class="button is-rounded" v-on:click="dropped = !dropped">
            <span class="icon is-small">
              <i class="fas fa-chevron-circle-down"></i>
            </span>
          </a>
          <a v-if="dropped" class="button is-rounded" v-on:click="dropped = !dropped">
            <span class="icon is-small">
              <i class="fas fa-chevron-circle-up"></i>
            </span>
          </a>
        </div>
      </div>
    </nav>
    <transition name="expand">
      <div v-if="dropped">
        <ul>
          <li>
            <nav class="level">
              <div class="level-left">
                <div class="level-ite">
                  <p class="menu-label">Under 10</p>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <b-switch size="is-small" :value="true" type="is-success"></b-switch>
                </div>
              </div>
            </nav>
          </li>
          <li>
              <nav class="level">
              <div class="level-left">
                <div class="level-ite">
                  <p class="menu-label">10-17</p>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <b-switch size="is-small" :value="true" type="is-success"></b-switch>
                </div>
              </div>
            </nav>
          </li>
          <li>
              <nav class="level">
              <div class="level-left">
                <div class="level-ite">
                  <p class="menu-label">18-24</p>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <b-switch size="is-small" :value="true" type="is-success"></b-switch>
                </div>
              </div>
            </nav>
          </li>
        </ul>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dropped: false
    };
  }
};
</script>

<style css scoped>
li {
  opacity: 1;
  transition: opacity 0.75s;
}
/* USE THIS WHEN ALL LISTS ARE AUTO GENERATED */
/* li:first-child {
  border-top: 1px solid #ddd;
} */ 
li {
  border-top: 1px solid #ddd;
} 
.box {
  margin-top: 20px;
}
.expand-enter {
  max-height: 0;
  margin: 0;
}
.expand-enter-to {
  max-height: 445px;
  transition: max-height 0.75s;
  margin: 0 auto 16px auto;
}
.expand-leave {
  max-height: 445px;
  margin: 0 auto 16px auto;
}
.expand-leave-to {
  max-height: 0;
  transition: max-height 0.75s;
  margin: 0;
}
</style>
