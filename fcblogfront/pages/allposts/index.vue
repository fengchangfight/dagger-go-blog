<template>
  <div class="allposts-page">
      <el-dialog
      title="写博文"
      :visible.sync="addPostDialogVisible"
      width="100%"
      :fullscreen="true"
      :before-close="handleClose">
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="标题">
          <el-input v-model="form.title" maxlength="300"></el-input>
        </el-form-item>
        <el-form-item label="正文">
          <wysiwyg v-model="form.content" />
        </el-form-item>
         <el-form-item label="分类">
              <el-select v-model="form.category_id" placeholder="文章分类">
    <el-option
      v-for="item in available_categories"
      :key="item.id"
      :label="item.name"
      :value="item.id">
    </el-option>
  </el-select>

          
        </el-form-item>

      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="addPostDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addBlogDoStuff">创建</el-button>
      </span>
    </el-dialog>

    <div style="width:100%;height:100%;">
          <app-header></app-header>

          <div class="post-body-center">
             <button v-if="whoami!=null && whoami.length>0" @click="addBlog" title="添加博文" class="add-fcblog"></button>
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
import {AXIOS} from '~/common/http-commons'
import base from '~/mixins/base'
import '~/css/dagger-blog.css'
import swal from 'sweetalert'
export default {
  components: {
      'app-header': Header
  },
  mixins:[base],
  data(){
    return{
        form:{
            title:'',
            content:'',
            category_id:''
        },
        available_categories:[],
        addPostDialogVisible: false,
        totalCount: 0,
        currentPage: 1,
        pageSize: 0,
        post_list:[]
    }
  },
  methods:{
      
      
      goPost(id){
          this.$router.push('/post/'+id);
      },
      loadCategories(){
          AXIOS.get('/api/v1/category').then(response=>{
              this.available_categories = response.data.data;
          }).catch(e=>{

          })
      },
      handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
      },
      addBlogDoStuff(){
          if(!this.checkNotEmpty(this.form.title)){
              swal ( "提示" ,  "标题不能为空" ,  "info" );
              return;
          }
          
          if(!this.checkNotEmpty(this.form.category_id)){
              swal ( "提示" ,  "分类不能为空" ,  "info" );
              return;
          }
          AXIOS.post('/api/v1/post',{
              title: this.form.title,
              content: this.form.content,
              category_id: this.form.category_id
          }).then(response=>{
              if(response.data.ok==true){
                  this.$notify({
                          title: '成功',
                          type: 'success',
                          message: response.data.message
                        });
                  this.addPostDialogVisible=false;
                  this.loadPosts();
              }else{
                  this.$notify.error({
                          title: '错误',
                          message: response.data.message
                        });
              }
          }).catch(e=>{
               this.$notify.error({
                          title: '错误',
                          message: '未知错误'
                        });
          })
      },
      addBlog(){
          this.addPostDialogVisible=true;
          this.form={};
          this.loadCategories();
      },
      changePage: function (currentPage) {
       this.currentPage = currentPage;
       this.loadPosts();
      },
      loadPosts(){
       var params = {};
       params.page = this.currentPage;
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
      this.whoamifoo();
      this.loadPosts();
  }
}
</script>

<style scoped>
.allposts-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;
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

.nav-container{
  max-width: 1000px;
  margin: 0 auto;
  margin-top: 10px;
}

nav a:hover{
  background: rgba(58, 42, 42, 0.6);
}

.post-body-center{
    margin: 0 auto; 
    min-width: 900px; 
    width:900px;
    min-height:700px;
    display: flex;
    flex-direction: column;
    margin-top:10px;
    position: relative;
}
.pagi-container{
    margin: 0 auto; 
    min-width: 700px; 
    width:700px;
    height: 10px;
    margin-top: 15px;
}



</style>

