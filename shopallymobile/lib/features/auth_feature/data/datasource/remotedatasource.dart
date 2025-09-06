import 'package:connectivity_plus/connectivity_plus.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:shopallymobile/features/auth_feature/data/model/user_model.dart';

class SignInCancelledException implements Exception {
  final String message;
  SignInCancelledException(this.message);
  
  @override
  String toString() => message;
}

abstract class RemoteDataSource {
  Future<UserModel> signinWithGoogle();
  Future<void> signout();
}

class RemoteDataSourceImpl implements RemoteDataSource {
  static final GoogleSignIn _googleSignIn = GoogleSignIn(
    scopes: ['email', 'profile'],
  );
  Future<bool> hasInternetConnection() async {
    final connectionResult = await Connectivity().checkConnectivity();
    if (connectionResult == ConnectivityResult.none) {
      return false;
    }
    return true;
  }

  @override
  Future<void> signout() async {
    try {
      await _googleSignIn.signOut();
    } catch (e) {
      throw e.toString();
    }
  }

  @override
  Future<UserModel> signinWithGoogle() async {
    final hasConnection = await hasInternetConnection();
    if (!hasConnection) {
      throw Exception('No internet connection');
    }
    try {
      await _googleSignIn.signOut();
    } catch (e) {
      // Ignore sign out errors
    }

    try {
      final GoogleSignInAccount? account = await _googleSignIn.signIn();
      if (account == null) {
        // User cancelled sign-in, return null or throw a specific cancellation exception
        throw SignInCancelledException('Sign-in was cancelled');
      }

      await account.authentication;
      return UserModel(
        name: account.displayName ?? '',
        email: account.email,
        photourl: account.photoUrl,
      );
    } catch (e) {
      if (e is SignInCancelledException) {
        rethrow;
      }
      throw Exception('Sign-in failed: ${e.toString()}');
    }
  }
}
