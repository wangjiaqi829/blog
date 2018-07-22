package blog

type MainController struct {
	BaseController
}

//首页, 只显示前N条
func (this *MainController) Index() {
	/*var list []*models.Article
	query := new(models.Article).Query()
	count, _ := query.Count()
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize), "/index%d.html").ToString()*/
	this.display("index")
}



//留404页面
func (this *MainController) Go404() {
	this.display("404")
}


//文章显示
func (this *MainController) Show() {
	/*var (
		post *models.Post = new(models.Post)
		err  error
	)
	urlname := this.Ctx.Input.Param(":urlname")
	if urlname != "" {
		post.Urlname = urlname
		err = post.Read("urlname")
	} else {
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		post.Id = int64(id)
		err = post.Read()
	}
	if err != nil || post.Status != 0 {
		this.Redirect("/404.html", 302)
	}
	post.Views++
	post.Update("Views")
	models.Cache.Delete("hotblog")
	post.Content = strings.Replace(post.Content, "_ueditor_page_break_tag_", "", -1)
	pre, next := post.GetPreAndNext()
	this.Data["post"] = post
	this.Data["pre"] = pre
	this.Data["next"] = next
	this.Data["smalltitle"] = "文章内容"
	if urlname == "about.html" {
		this.Data["smalltitle"] = "关于我"
	}

	this.setHeadMetas(post.Title, strings.Trim(post.Tags, ","), post.Title)
	this.display("article")*/
}
