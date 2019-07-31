import 'entry.dart';
import 'package:dio/dio.dart';

Future<List<Data>> getBlogs() async {
  Response response;
  Dio dio = new Dio();
  response = await dio.get("http://127.0.0.1:8080/api/v1/blog/search");

  if (response.statusCode == 200) {
    BlogResponse result = new BlogResponse.fromJson(response.data);
    if (result.code == 0) {
      return result.data;
    }
  }

  return [];
}
