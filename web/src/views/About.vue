<template>
  <section class="section">
    <div class="columns">
      <div class="column is-6 is-offset-3">
        <div class="content is-medium">
          <h1>About</h1>
          <p>
            This website is aimed at making stop and search data from England, Wales and
            Northern Ireland easier to explore and understand. The website is in the
            early stages of development with many more features planned.
          </p>
          <h3>How can I trust that you are displaying accurate data?</h3>
          <p>
            You can't. That is why all the source code for this project is open sourced
            at
            <a href="https://github.com/tjcain/stopandsearch">github</a>.
          </p>
          <h3>Where does the information come from?</h3>
          <p>
            The data displayed all comes from the police's open data about crime and policing
            in England, Wales and Northern Ireland, which can be found
            <a
              href="https://data.police.uk"
            >here</a>.
          </p>
          <h3>Why can I only see data for West Midlands Police?</h3>
          <p>
            The website is still in the early stages of development. It is online to gather
            feedback from those who may be interested in using it. If you would be interested
            in seeing the information for your police force displayed please
            <router-link to="/contact">contact</router-link> me.
          </p>
          <h3>Something broke, or doesn't work propely!</h3>
          <p>
            Unforunately there may be bugs, if you find one please consider
            <router-link to="/contact">letting me know</router-link>, or
            create an issue on
            <a href="https://github.com/tjcain/stopandsearch/issues">github</a>.
          </p>
          <h3>Why does the website look so bad?</h3>
          <p>
            That happens when I am left to design something by myself.. If you
            have any suggestions you could consider
            <router-link to="/contact">letting me know</router-link>, or
            create an issue on
            <a href="https://github.com/tjcain/stopandsearch/issues">github</a>.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.content {
  min-height: 70vh;
}
</style>

