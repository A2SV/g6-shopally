import 'package:connectivity_plus/connectivity_plus.dart';

abstract class NetworkInfo {
  Future<bool> get isConnected;
}

class NetworkInfoImpl implements NetworkInfo {
  final Connectivity connectivity;

  NetworkInfoImpl(this.connectivity);

  @override
  Future<bool> get isConnected async {
    final result = await connectivity.checkConnectivity();
    // The connectivity result is a list now, so we check if it's not empty and contains a valid connection type.
    return result.contains(ConnectivityResult.mobile) ||
        result.contains(ConnectivityResult.wifi);
  }
}
