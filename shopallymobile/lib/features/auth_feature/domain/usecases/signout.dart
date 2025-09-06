import 'package:shopallymobile/features/auth_feature/domain/repositories/user_repo.dart';

class Signout {
  final UserRepository userRepository;

  Signout(this.userRepository);

  Future<void> call() async {
    return await userRepository.signout();
  }
}
