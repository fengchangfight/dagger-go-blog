<template>
  <div class="picture-page">
    <div style="width:100%;height:100%;">
          <app-header></app-header>
          <div class="back-music-fcblog">
                   <audio style="width: 320px;height:40px;" src="http://www.???.cc/fcnotes/EpicHollywoodTrailer.mp3" controls autoplay loop>
                   </audio>
                </div>
          <div class="gallery-container">
            <gallery :images="images" :index="index" @close="index = null"></gallery>
              <div
                class="image"
                v-for="(image, imageIndex) in images"
                :key="imageIndex"
                @click="index = imageIndex"
                :style="{ backgroundImage: 'url(' + image + ')', width: '300px', height: '200px' }"
               ></div>
          </div>
      </div>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import '~/css/dagger-blog.css'

export default {
  components: {
      'app-header': Header
  },
  data(){
    return{
images: [
          'http://www.???.???/imageprocessing/fit?width=1920&height=1080&type=jpeg&file=3RrkufxmI78ugSVR.jpeg',
          'http://www.????.??/imageprocessing/fit?width=1920&height=1080&type=jpeg&file=bftZJHMRbXhjbeXx.jpeg',
          'http://www.???/imageprocessing/fit?width=1920&height=1080&type=jpeg&file=Q5Lt529Ri5UQYUQr.jpeg',
          'http://www.???.???/imageprocessing/fit?width=1920&height=1080&type=jpeg&file=LFSuPgAwp93DwwR2.jpeg',
          'http://www.???.???/imageprocessing/fit?width=1920&height=1080&type=jpeg&file=d63hhAwZmTcfemzm.jpeg',
          'http://www.???.???/imageprocessing/fit?width=1920&height=1080&type=jpeg&file=FsJcd5ehSaAB9x5N.jpeg'
        ],
        index: null
    }
  }
}
</script>

<style scoped>
.picture-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;

}

.image {
    float: left;
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center center;
    border: 1px solid #ebebeb;
    margin: 5px;
}
.gallery-container{
  width: 950px;
  min-width: 800px;
  min-height:700px;
  overflow: scroll;
  margin: 0 auto;
  margin-top: 2px;
}

.back-music-fcblog{
  margin-top: 30px;
}

</style>
