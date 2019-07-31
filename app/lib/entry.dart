class BlogResponse {
  int code;
  String message;
  List<Data> data;

  BlogResponse({this.code, this.message, this.data});

  BlogResponse.fromJson(Map<String, dynamic> json) {
    code = json['Code'];
    message = json['Message'];
    if (json['Data'] != null) {
      data = new List<Data>();
      json['Data'].forEach((v) {
        data.add(new Data.fromJson(v));
      });
    }
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['Code'] = this.code;
    data['Message'] = this.message;
    if (this.data != null) {
      data['Data'] = this.data.map((v) => v.toJson()).toList();
    }
    return data;
  }
}

class Data {
  String iD;
  String title;
  String summary;
  String published;
  String updated;
  Author author;
  Link link;
  String diggs;
  String views;
  String comments;

  Data(
      {this.iD,
      this.title,
      this.summary,
      this.published,
      this.updated,
      this.author,
      this.link,
      this.diggs,
      this.views,
      this.comments});

  Data.fromJson(Map<String, dynamic> json) {
    iD = json['ID'];
    title = json['Title'];
    summary = json['Summary'];
    published = json['Published'];
    updated = json['Updated'];
    author =
        json['Author'] != null ? new Author.fromJson(json['Author']) : null;
    link = json['Link'] != null ? new Link.fromJson(json['Link']) : null;
    diggs = json['Diggs'];
    views = json['Views'];
    comments = json['Comments'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['ID'] = this.iD;
    data['Title'] = this.title;
    data['Summary'] = this.summary;
    data['Published'] = this.published;
    data['Updated'] = this.updated;
    if (this.author != null) {
      data['Author'] = this.author.toJson();
    }
    if (this.link != null) {
      data['Link'] = this.link.toJson();
    }
    data['Diggs'] = this.diggs;
    data['Views'] = this.views;
    data['Comments'] = this.comments;
    return data;
  }
}

class Author {
  String name;
  String uRI;
  String avatar;

  Author({this.name, this.uRI, this.avatar});

  Author.fromJson(Map<String, dynamic> json) {
    name = json['Name'];
    uRI = json['URI'];
    avatar = json['Avatar'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['Name'] = this.name;
    data['URI'] = this.uRI;
    data['Avatar'] = this.avatar;
    return data;
  }
}

class Link {
  String rel;
  String href;

  Link({this.rel, this.href});

  Link.fromJson(Map<String, dynamic> json) {
    rel = json['Rel'];
    href = json['Href'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['Rel'] = this.rel;
    data['Href'] = this.href;
    return data;
  }
}
