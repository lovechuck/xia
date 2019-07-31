import 'package:url_launcher/url_launcher.dart';

import 'entry.dart';
import 'package:flutter/material.dart';

class CardTitle extends StatelessWidget {
  final int index;
  final Data data;

  CardTitle({this.index, this.data});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      contentPadding: EdgeInsets.all(10.0),
      title: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Container(
            padding: EdgeInsets.only(
              bottom: 12.0,
            ),
            child: InkWell(
              // When the user taps the button, show a snackbar.
              onTap: () {
                _launchURL(data.link.href);
              },
              child: Text(
                data.title,
                style: TextStyle(
                  fontStyle: FontStyle.italic,
                ),
              ),
            ),
          ),
        ],
      ),
      subtitle: Text(data.summary),
    );
  }
}

_launchURL(url) async {
  if (await canLaunch(url)) {
    await launch(url);
  } else {
    throw 'Could not launch $url';
  }
}
