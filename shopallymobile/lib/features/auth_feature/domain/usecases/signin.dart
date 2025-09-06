import 'package:shopallymobile/features/auth_feature/domain/entities/user.dart';
import 'package:shopallymobile/features/auth_feature/domain/repositories/user_repo.dart';

class Signin {
  final UserRepository userRepository;

  Signin(this.userRepository);

  Future<AuthUser?> call() async {
    return await userRepository.signinWithGoogle();
  }
}
