<template>
  <div class="life-page">
    <div style="width:100%;height:100%;">
          <app-header></app-header>
          <div class="post-body-center">
             <div class="card" v-bind:key="item.id" v-for="item in post_list">
                 <div @click="goPost(item.id)"  style="float:left;"><img width="90" height="90" :src="'imgs/'+head_img"/></div>
                 <div class="content-title">
                     <p @click="goPost(item.id)"  class="title-text"><label class="title-label">{{item.title}}</label></p>
                     <p class="meta-info">发布于：{{item.created_at}}  &nbsp;&nbsp; 类别：{{item.category}}</p>
                 </div>
                 <button v-if="whoami!=null && whoami.length>0" @click="trashBlog(item.id)" title="移到垃圾箱" class="trash-fcblog"></button>
                 <button v-if="whoami!=null && whoami.length>0" @click="editBlog(item.id)" title="编辑文章" class="edit-fcblog"></button>
             </div>
            
             <div v-if="totalCount>0" class="pagi-container">
               <el-pagination style="margin: 0 auto" class="pagination" layout="prev, pager, next" :total="totalCount" :page-size="pageSize"
               v-on:current-change="changePage">
                </el-pagination>
              </div>
          </div>
      </div>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import base from '~/mixins/base'
import {AXIOS} from '~/common/http-commons'
import '~/css/dagger-blog.css'
export default {
  components: {
      'app-header': Header
  },
  mixins:[base],
  head () {
    return {
      title: '生活篇',
      meta: [
        { hid: 'Description', name: 'Description', content: '生活篇' },
        { hid: 'Keywords', name: 'Keywords', content: '生活篇' }
      ]
    }
  },
  data(){
    return{
        post_list:[],
        totalCount:0,
        currentPage:1,
        pageSize:0

    }
  },
  methods:{
      changePage: function (currentPage) {
         this.currentPage = currentPage;
         this.loadPostsCateLife();
      },
      loadPostsCateLife(){
          var params = {};
          params.page = this.currentPage;
          params.category_id = 1;
          AXIOS.get('/public/blog',{
            params: params
          }).then(response=>{
              this.post_list = response.data.data.data;
              this.totalCount = Number(response.data.data.totalCount);
              this.currentPage = Number(response.data.data.currentPage);
              this.pageSize = Number(response.data.data.pageSize);
          }).catch(e=>{
            console.log(e);
          })
      }

  },
  mounted(){
      this.loadPostsCateLife();
      this.whoamifoo();
  }
}
</script>

<style scoped>
.life-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;

}
nav{
  background: rgba(70, 66, 66, 0.5);
  box-shadow: 0px 7px 7px rgb(61, 58, 58);
  overflow: auto;
  width: 100;
  border-radius: 15px 50px 30px;
}
ul{
  list-style: none;
  margin: 0;
  padding: 0;
}
li{
  float: left;
}
nav a{
  width: 120px;
  display: block;
  padding: 15px 15px;
  font-family: HYDaSongJ;
  font-size: 17px;
  color: white;
  text-align: center;
  text-decoration: none;
}
.body-container-cover{
  width: 100%;
  height: 100%;
  background-color: black;
}
.body-container{
  width: 100%;
  height: 100%;
  background-color: white;
}

.mainimg:hover{
  cursor: pointer;
  opacity: 0.8;
}
.bodyimg:hover{

}

.nav-container{
  max-width: 1000px;
  margin: 0 auto;
  margin-top: 10px;
}

nav a:hover{
  background: rgba(58, 42, 42, 0.6);
}

</style>

