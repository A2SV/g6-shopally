class AuthUser {
  final String name;
  final String email;
  String? photourl;
  String? language;
  String? currency;
  bool? notifications;
  

   AuthUser({
    required this.name,
    required this.email,
    this.photourl,
    this.language,
    this.currency,
    this.notifications,
  
  });
}
