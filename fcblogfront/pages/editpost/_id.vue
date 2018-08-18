<template>
  <div class="edit-post-page">
    <div style="width:100%;height:100%;">
          <app-header></app-header>
          <div class="post-body">
            <div style="display:flex;margin-top:30px;">
                <div style="height:30px;width:20%;">
                <el-select v-model="form.category_id" placeholder="文章分类">
                  <el-option
                    v-for="item in available_categories"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                  </el-option>
                </el-select>
                </div>
                <div style="flex:1"><el-input v-model="form.title"></el-input></div>
                
            </div>
            
            
            <div>
                <wysiwyg v-model="form.content" />
            </div>
            <div style="display:block">
              <el-button style="float:right;margin-left:10px;" type="primary" @click="editBlogDoStuff">更新</el-button>
              <el-button style="float:right" type="success" @click="goPage('/post/'+current_blog_id)">进入博文</el-button>
           </div>
          </div>
      </div>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import {AXIOS} from '~/common/http-commons'
import '~/css/dagger-blog.css'
import swal from 'sweetalert'
import base from '~/mixins/base'
export default {
  components: {
      'app-header': Header
  },
  data(){
    return{
        current_blog_id:'',
        postDetail:{
                        
        },
        available_categories:[],
        form:{}
    }
  },
  mixins:[base],
  methods:{
      editBlogDoStuff(){
          if(!this.checkNotEmpty(this.form.title)){
              swal ( "提示" ,  "标题不能为空" ,  "info" );
              return;
          }
          
          if(!this.checkNotEmpty(this.form.category_id)){
              swal ( "提示" ,  "分类不能为空" ,  "info" );
              return;
          }
          AXIOS.put('/api/v1/post/'+this.current_blog_id,{
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
                  this.loadPostDetail(this.current_blog_id);
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
      loadCategories(){
          AXIOS.get('/api/v1/category').then(response=>{
              this.available_categories = response.data.data;
          }).catch(e=>{

          })
      },
      loadPostDetail(id){
          AXIOS.get('/public/post/'+id).then(response=>{
              if(response.data.ok==true){

                  this.form=response.data.data;
                  console.log(this.form);
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
      }
  },
  
  mounted(){
      this.current_blog_id=this.$route.params.id;
      this.loadPostDetail(this.current_blog_id);
      this.loadCategories();
  }
}
</script>

<style scoped>
.edit-post-page {
  min-height: 100vh;
  width: 100%;
  height: 100%;
  display: flex;

}

.post-body{
    width: 1000px;
    min-width: 800px;
    min-height: 700px;
    margin: 0 auto;
}
.meta-info{
    font-size: 70%;
    color: gray;
    float: right;
    margin-right:20px;
    margin-top: 2px;
}

h3{
      justify-content: center;
      align-items: center;
      text-align: center;
}

</style>

