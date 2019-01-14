<template>
  <section class="section">
    <div class="columns">
      <div class="column is-6 is-offset-1">
        <div class="content is-medium">
          <h1>Sorry, something went wrong</h1>
          <p>
            This error has been logged and will be looked into, please
            <router-link to="/">try again</router-link> or if the error
            persists consider <router-link to="/contact">letting me know</router-link>.
          </p>
          
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.content {
  min-height: 60vh;
}
</style>

