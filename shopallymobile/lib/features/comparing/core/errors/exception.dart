class ServerException implements Exception {
  final String message;
  ServerException({required this.message});

  @override
  String toString() => message;
}

class CacheException implements Exception {}

class NetworkException implements Exception {}

class NotFoundException implements Exception {
  final String message;

  NotFoundException(this.message);
}