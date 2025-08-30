import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/auth_feature/presentation/bloc/bloc.dart';
import 'package:shopallymobile/auth_feature/presentation/bloc/event.dart';
import 'package:shopallymobile/auth_feature/presentation/bloc/state.dart';

import 'package:shopallymobile/auth_feature/presentation/pages/widgets.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key, required this.userRepository});
  final UserRepository userRepository;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) =>
          UserAuthBloc(userRepository)..add(GetAuthenticatedUserEvent()),
      child: Scaffold(
        backgroundColor: const Color(0xFFF4F5F7),
        body: SafeArea(
          child: BlocBuilder<UserAuthBloc, UserAuthState>(
            builder: (context, state) {
              if (state is LoadingState) {
                return const Center(child: CircularProgressIndicator());
              }
              if (state is ErrorState) {
                return Center(child: Text(state.message));
              }
              final user = state is SuccessState ? state.user : null;
              print('++++++ user saved++++++');
              
              return SingleChildScrollView(
                padding: const EdgeInsets.symmetric(
                  horizontal: 16,
                  vertical: 12,
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.stretch,
                  children: [
                    // Header
                    Column(
                      children: [
                        const SizedBox(height: 70),
                        _Avatar(
                          name: user?.name,
                          photoUrl: user?.photourl,
                          fallbackInitial: _initialFromId('r'),
                        ),
                        const SizedBox(height: 12),
                        Text(
                          (user != null && user.name.isNotEmpty)
                              ? user.name
                              : 'Guest',
                          style: Theme.of(context).textTheme.titleLarge
                              ?.copyWith(fontWeight: FontWeight.w600),
                        ),
                        const SizedBox(height: 6),
                        Text(
                          (user != null && user.email.isNotEmpty)
                              ? user.email
                              : 'Not signed in',
                          style: Theme.of(context).textTheme.bodyMedium
                              ?.copyWith(color: Colors.black54),
                        ),
                      ],
                    ),
                    const SizedBox(height: 12),
                    // Quick actions
                    Container(
                      margin: const EdgeInsets.only(left: 20),
                      child: Row(
                        children: [
                          QuickTile(
                            icon: Icons.favorite_border,
                            label: 'Favorites',
                            onTap: () {},
                          ),
                          const SizedBox(width: 10),
                          QuickTile(
                            icon: Icons.notifications_none,
                            label: 'Notifications',
                            onTap: () {},
                          ),
                        ],
                      ),
                    ),
                    const SizedBox(height: 12),
                    // Settings card (static UI labels to match appearance)
                    Container(
                      decoration: _cardDecoration(),
                      child: Column(
                        children: [
                          SettingsRow(
                                title: 'Language',
                                trailingText: user?.language ?? 'English',
                                onTap: () async {
                                  if (user == null) {
                                    showSocialLoginBottomSheet(
                                      context,
                                      onGoogle: () => context.read<UserAuthBloc>().add(SignInEvent()),
                                    );
                                    return;
                                  }
                                  final selected = await showModalBottomSheet<String>(
                                    context: context,
                                    builder: (ctx) {
                                      final options = const ['English', 'Amharic', ];
                                      return SizedBox(
                                        height: 200,
                                        child: Padding(
                                          padding: const EdgeInsets.only(left: 20 ,top: 20),
                                          child: ListView(
                                            children: options.map((lang) {
                                              return ListTile(
                                                title: Text(lang),
                                                trailing: (user?.language ?? 'English') == lang
                                                    ? const Icon(Icons.check)
                                                    : null,
                                                onTap: () => Navigator.pop(ctx, lang),
                                              );
                                            }).toList(),
                                          ),
                                        ),
                                      );
                                    },
                                  );
                                  if (selected != null && selected.isNotEmpty) {
                                    context.read<UserAuthBloc>().add(UpdateLanguageEvent(selected));
                                  }
                                },
                              ),
                          const Divider(height: 1),
                          SettingsRow(
                            title: 'Currency',
                            trailingText: user?.currency ?? 'USD',
                            onTap: () async {
                              if (user == null) {
                                showSocialLoginBottomSheet(
                                  context,
                                  onGoogle: () => context
                                      .read<UserAuthBloc>()
                                      .add(SignInEvent()),
                                );
                                return;
                              }
                              final selected = await showModalBottomSheet<String>(
                                context: context,
                                builder: (ctx) {
                                  final options = const ['USD',  'BIRR', ];
                                  return SizedBox(
                                    height: 200,
                                    child: Padding(
                                      padding:EdgeInsetsGeometry.only(top:20 , left:20) , 
                                      child: ListView(
                                        children: options.map((c) {
                                          return ListTile(
                                            title: Text(c),
                                            trailing: (user?.currency ?? 'USD') == c
                                                ? const Icon(Icons.check)
                                                : null,
                                            onTap: () => Navigator.pop(ctx, c),
                                          );
                                        }).toList(),
                                      ),
                                    ),
                                  );
                                },
                              );
                              if (selected != null && selected.isNotEmpty) {
                                context
                                    .read<UserAuthBloc>()
                                    .add(UpdateCurrencyEvent(selected));
                              }
                            },
                          ),
                          const Divider(height: 1),
                          SwitchRow(
                            title: 'Notifications',
                            value: user?.notifications ?? true,
                            onChanged: (val) {
                              if (user == null) {
                                showSocialLoginBottomSheet(
                                  context,
                                  onGoogle: () => context
                                      .read<UserAuthBloc>()
                                      .add(SignInEvent()),
                                );
                                return;
                              }
                              context
                                  .read<UserAuthBloc>()
                                  .add(UpdateNotificationEvent(val));
                            },
                          ),
                          const Divider(height: 1),
                        ],
                      ),
                    ),
                    const SizedBox(height: 12),
                    // Version card
                    Container(
                      decoration: _cardDecoration(),
                      child: const ListTile(
                        title: Text('Version'),
                        subtitle: Text('1.1.7'),
                        trailing: Text(
                          'UPDATE',
                          style: TextStyle(color: Colors.black),
                        ),
                      ),
                    ),
                    const SizedBox(height: 12),
                    // Sign in / out action (UI only; uses existing bloc events)
                    Container(
                      decoration: _cardDecoration(),
                      child: ListTile(
                        leading: Icon(
                          user != null ? Icons.logout : Icons.login,
                          color: Colors.redAccent,
                        ),
                        title: Text(
                          user != null ? 'Sign out' : 'Sign in',
                          style: const TextStyle(
                            color: Colors.redAccent,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                        trailing: const Icon(
                          Icons.chevron_right,
                          color: Colors.redAccent,
                        ),
                        onTap: () {
                          if (user == null) {
                              

                            
                            showSocialLoginBottomSheet(
                              context,
                              onGoogle: () => context.read<UserAuthBloc>().add(
                                SignInEvent(),
                              ),
                            );
                          } else {
                            showSignOutDialog(context).then((ok) {
                              if (ok == true) {
                                context.read<UserAuthBloc>().add(
                                  SignOutEvent(),
                                );
                              }
                            });
                          }
                        },
                      ),
                    ),
                  ],
                ),
              );
            },
          ),
        ),
      ),
    );
  }
}

BoxDecoration _cardDecoration() {
  return BoxDecoration(
    color: Colors.white,
    borderRadius: BorderRadius.circular(16),
    boxShadow: const [
      BoxShadow(color: Color(0x11000000), blurRadius: 8, offset: Offset(0, 2)),
    ],
  );
}

String _initialFromId(String id) {
  final letters = id.trim();
  if (letters.isEmpty) return 'R';
  return letters.substring(0, 1).toUpperCase();
}

void _noopToggle(bool _) {}

class _Avatar extends StatelessWidget {
  final String? name;
  final String? photoUrl;
  final String fallbackInitial;
  const _Avatar({
    required this.name,
    required this.photoUrl,
    required this.fallbackInitial,
  });

  @override
  Widget build(BuildContext context) {
    if (photoUrl != null && photoUrl!.isNotEmpty) {
      return CircleAvatar(radius: 48, backgroundImage: NetworkImage(photoUrl!));
    }
    final initial = (name != null && name!.isNotEmpty)
        ? name!.trim().substring(0, 1).toUpperCase()
        : fallbackInitial;
    return CircleAvatar(
      radius: 48,
      backgroundColor: const Color(0xFF27C08A),
      child: Text(
        initial,
        style: const TextStyle(
          color: Colors.white,
          fontSize: 24,
          fontWeight: FontWeight.w700,
        ),
      ),
    );
  }
}
