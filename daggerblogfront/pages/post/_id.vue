<template>
  <div class="post-page">
    <div style="width:100%;height:100%;">
          <app-header></app-header>

          <div class="post-body" v-loading="loading_content">
            <div style="position:relative;">
                <button v-if="whoami!=null && whoami.length>0" @click="editBlog" title="编辑文章" class="edit-fcblog"></button>
                <h3 style="margin-left:25px;margin-right:280px;">{{postDetail.title}}</h3>
                <div class="back-music-fcblog">
                   <audio style="width: 280px;height:40px;" src="http://www.???.cc/fcnotes/empireofangels.mp3" controls autoplay loop>
                   </audio>
                </div>
            </div>
            <div style="height:30px;width:100%;margin-top:50px;"><p class="meta-info-view">发布于：{{postDetail.created_at}}&nbsp;&nbsp; 类别：{{postDetail.category}}</p></div>
            <div class="body-content"  v-html="postDetail.content"></div>
            <div class="comment-section">
                <div class="give-comment">
                    <div class="comment-meta">
                        <div class="pinglunicon"></div>
                        <el-input v-model="commenter" placeholder="评前留名"></el-input>
                        <!-- <drag-verify style="margin-top:10px;" :width="240"
			              :height="30"
			              :text="dragText"
			              :success-text="successText"
                          @passcallback="passDrag" >
                        </drag-verify> -->
                        <el-switch
                           v-model="verified"
                           active-text="验证"
                           inactive-text="取消"
                           @change="passDrag">
                       </el-switch>

                    </div>
                    <div class="comment-content">
                        <el-input
                          type="textarea"
                          :rows="5"
                          placeholder="想说的话..."
                          v-model="usercomment">
                        </el-input>
                        <el-button @click="doComment" style="position:absolute;left:680px;top:78px;" type="primary">评论</el-button>
                    </div>
                </div>
                <div class="comment-list">
                  <div v-for="v in comment_list" v-bind:key="v.id" class="comment-card">
                   <hr width="900">
                   <label style="opacity:0.7;font-size:70%;"><b>{{v.commenter}}</b>&nbsp;于&nbsp;{{v.comment_time}}&nbsp;评论道:</label>
                   <p>{{v.comment_content}}</p>
                  </div>
                </div>
                <div class="bottom-place-taker">
                </div>

            </div>

          </div>
      </div>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import {AXIOS} from '~/common/http-commons'
import '~/css/dagger-blog.css'
import base from '~/mixins/base'
import swal from 'sweetalert'
export default {
  components: {
      'app-header': Header
  },
  asyncData ({ params }) {
      return AXIOS.get(`/public/post/${params.id}`).then(response=>{
            if(response.data.ok==true){
              return { title: response.data.data.title }
            }else{
            }
          }).catch(e=>{

          })
  },
  mixins:[base],
  head () {
    return {
      title: this.title,
      meta: [
        { hid: 'Description', name: 'Description', content: '阿昌的博客' },
        { hid: 'Keywords', name: 'Keywords', content: this.title }
      ]
    }
  },
  methods:{
      clearUnread4Post(){
          AXIOS.put('/api/v1/comment/readcomment/'+this.current_blog_id).then(response=>{
              if(response.data.ok==true){
                  console.log(response.data)
              }
          }).catch(e=>{

          })
      },

      loadComment(){
          AXIOS.get('/public/comment4blog/'+this.current_blog_id).then(response=>{
              if(response.data.ok==true){
                  this.comment_list = response.data.data;
              }else{
                   this.$notify.error({
                          title: '错误',
                          message: response.data.message
                        });
              }
          }).catch(e=>{
               this.$notify.error({
                          title: '错误',
                          message: "载入评论未知错误"
                        });
          })
      },
      doComment(){
          if(this.comment_time_window==false){
              swal ( "提示" ,  "没验证不能评论哦(づ￣ 3￣)づ" ,  "info" );
              return;
          }
          if(this.commenter!=null && this.commenter.length<1){
              swal ( "提示" ,  "不留个名就想评论？(づ￣ 3￣)づ" ,  "info" );
              return;
          }
          if(this.commenter.length>15){
              swal ( "提示" ,  "名字取太长了啊!(づ￣ 3￣)づ" ,  "info" );
              return;
          }
          if(this.usercomment!=null && this.usercomment.length<1){
              swal ( "提示" ,  "评论总得说点什么吧？(づ￣ 3￣)づ" ,  "info" );
              return;
          }

          AXIOS.post('/public/comment-stuff',{
              comment_content: this.usercomment,
              commenter: this.commenter,
              session_id: this.comment_session_id,
              blog_id: this.current_blog_id
          }).then(response=>{
              if(response.data.ok==true){
                 this.$notify({
                          title: '成功',
                          type: 'success',
                          message: response.data.message
                        });
                 this.loadComment();
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
      passDrag(){
          // generate rand string of 16 length
          if(this.verified==false){
              this.comment_time_window=false;
              return;
          }

          function makeSessionId() {
            var text = "";
            var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
            for (var i = 0; i < 16; i++){
            text += possible.charAt(Math.floor(Math.random() * possible.length));
            }
            return text;
          }

          this.comment_session_id = makeSessionId()
          // use that string as key for redis, value should be post id
          AXIOS.post('/public/comment-2min-window',{
              sessionId : this.comment_session_id,
              blogId:this.current_blog_id
          }).then(response=>{
              if(response.data.ok==true){
                  console.log("Set comment window for 2 mins with sessionId:"+comment_session_id)
              }else{
                  console.log("comment window open failed...")
              }
          }).catch(e=>{
              console.log("comment window open failed with unknow reason...")
          })
          // expire in 2 mins

          // set send comment_window open
          this.comment_time_window=true;

      },
      saySuccess(){
          console.log('success')
      },
      sayFail(){
          console.log('fail')
      },
      editBlog(){
          this.$router.push('/editpost/'+this.current_blog_id);
      },
      loadPostDetail(id){
          this.loading_content=true;
          AXIOS.get('/public/post/'+id).then(response=>{
              if(response.data.ok==true){
                  this.postDetail=response.data.data;
                  this.loading_content=false;
              }else{
                  this.$notify.error({
                          title: '错误',
                          message: response.data.message
                        });
                this.$router.push('/allposts');
              }
          }).catch(e=>{
               this.$notify.error({
                          title: '错误',
                          message: '未知错误'
                        });
              this.$router.push('/allposts');
          })
      }
  },
  data(){
    return{
        verified: false,
        comment_list:[],
        commenter:'',
        comment_session_id: '',
        comment_time_window: false,
        dragText:'向右拖动验证',
        successText:'成功,两分钟内有效',
        loading_content: false,
        current_blog_id:'',
        usercomment:'',
        postDetail:{

        }
    }
  },
  mounted(){
      this.current_blog_id=this.$route.params.id;
      this.loadComment();
      this.loadPostDetail(this.current_blog_id);
      this.whoamifoo();
      this.clearUnread4Post();

  }
}
</script>

<style scoped>
.post-page {
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
.meta-info-view{
    font-size: 70%;
    float: right;
    margin-right:20px;
    margin-top: 2px;
}
.body-content{
    width: 1000px;
    margin: 0 auto;
    min-height:700px;
    text-indent: 10px;
    text-align: justify;
    overflow-x: scroll;
}
h3{
      justify-content: center;
      align-items: center;
      text-align: center;
}
.edit-fcblog{
    position: absolute;
    top: 2px;
    left: 6px;
    background:url('~/static/imgs/editbutton.png') no-repeat;
    background-size: 100% 100%;
    border:none;
    width:20px;
    height:20px;
}
.edit-fcblog:hover{
    cursor: pointer;
    opacity: 0.5;
}
.back-music-fcblog{
    position: absolute;
    top: 0px;
    right: 5px;
    height: 45px;
    /* background-color: yellow; */
    border:none;
}
.comment-section{
    width: 1000px;
    min-width: 1000px;
    margin: 0 auto;
    min-height:300px;
    position: relative;
}
.give-comment{
    width: 100%;
    height:140px;
    display: flex;
    flex-direction: row;

}
.comment-meta{
    width: 250px;
    height:100%;
}
.comment-content{
    flex:1;
    height:100%;
    position: relative;
}
.comment-section{

}
.pinglunicon{
    width: 40%;
    height:40px;
    background-color: bisque;
    margin: 0 auto;
    background:url('~/static/imgs/comment.png') no-repeat;
    background-size: 100% 100%;
}
.comment-list{
    width: 1000px;
    min-height:200px;
}
.comment-card{
    width: 900px;
    min-height:100px;
    /* background-color: yellow; */
    margin-top:10px;
    margin: 0 auto;
}
.bottom-place-taker{
    height:100px;
    width: 100%;
}
</style>
