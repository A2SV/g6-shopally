import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shopallymobile/features/auth_feature/domain/repositories/user_repo.dart';
import 'package:shopallymobile/features/auth_feature/presentation/bloc/bloc.dart';
import 'package:shopallymobile/features/auth_feature/presentation/bloc/event.dart';
import 'package:shopallymobile/features/auth_feature/presentation/bloc/state.dart';
import 'package:curved_navigation_bar/curved_navigation_bar.dart';

import 'package:shopallymobile/features/auth_feature/presentation/pages/widgets.dart';
import 'package:shopallymobile/core/localization/localization_store.dart';
import 'package:shopallymobile/core/localization/language_bloc.dart';
import 'package:shopallymobile/core/localization/language_event.dart';
import 'package:shopallymobile/core/localization/language_state.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({
    super.key,
    required this.userRepository,
    required this.isDark,
    required this.onDarkChanged,
  });
  final UserRepository userRepository;
  final bool isDark;
  final ValueChanged<bool> onDarkChanged;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) =>
          UserAuthBloc(userRepository)..add(GetAuthenticatedUserEvent()),
      child: BlocBuilder<LanguageBloc, LanguageState>(
        builder: (context, langState) {
          // Rebuild on language changes so getText(...) values update
          return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
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
                        avatar(
                          name: user?.name,
                          photoUrl: user?.photourl,
                          fallbackInitial: _initialFromId(''),
                        ),
                        const SizedBox(height: 12),
                        Text(
                          (user != null && user.name.isNotEmpty)
                              ? user.name
                              : getText('Guest'),
                          style: Theme.of(context).textTheme.titleLarge
                              ?.copyWith(fontWeight: FontWeight.w600),
                        ),
                        const SizedBox(height: 6),
                        Text(
                          (user != null && user.email.isNotEmpty)
                              ? user.email
                              : getText('not_signed_in'),
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
                            label: getText('favorites'),
                            onTap: () {},
                          ),
                          const SizedBox(width: 10),
                          QuickTile(
                            icon: Icons.notifications_none,
                            label: getText('notifications'),
                            onTap: () {},
                          ),
                        ],
                      ),
                    ),
                    const SizedBox(height: 12),
                    // Settings card (static UI labels to match appearance)
                    Container(
                      decoration: cardDecoration(context),
                      child: Column(
                        children: [
                          settingsRow(
                            title: getText('language' ,),
                            trailingText: user?.language ?? getText('english'),
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
                                  final options = [
                                    getText('english' , ),
                                    getText('amharic'),
                                  ];
                                  return SizedBox(
                                    height: 200,
                                    child: Padding(
                                      padding: const EdgeInsets.only(
                                        left: 20,
                                        top: 20,
                                      ),
                                      child: ListView(
                                        children: options.map((lang) {
                                          // Show a check based on the user's canonical saved name
                                          final currentLangName =
                                              (user.language ?? 'English')
                                                  .toLowerCase();
                                          final isChecked =
                                              (currentLangName == 'english' &&
                                                  lang == getText('english')) ||
                                              (currentLangName == 'amharic' &&
                                                  lang == getText('amharic'));
                                          return ListTile(
                                            title: Text(lang),
                                            trailing: isChecked
                                                ? const Icon(Icons.check)
                                                : null,
                                            onTap: () =>
                                                Navigator.pop(ctx, lang),
                                          );
                                        }).toList(),
                                      ),
                                    ),
                                  );
                                },
                              );
                              if (selected != null && selected.isNotEmpty) {
                                // Map localized label -> canonical name/code for persistence and localization
                                final isAmharic =
                                    selected.toLowerCase() ==
                                    getText('amharic').toLowerCase();
                                final code = isAmharic ? 'am' : 'en';
                                final canonicalName = isAmharic ? 'Amharic' : 'English';

                                // Update app localization
                                context.read<LanguageBloc>().add(
                                  SetLanguageEvent(code),
                                );

                                // Persist to user profile so trailing text updates
                                context.read<UserAuthBloc>().add(
                                  UpdateLanguageEvent(canonicalName),
                                );
                              }
                            },
                          ),
                          const Divider(height: 1),
                          settingsRow(
                            title: getText('currency'),
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
                                  final options = [
                                    getText('usd'),
                                    getText('birr'),
                                  ];
                                  return SizedBox(
                                    height: 200,
                                    child: Padding(
                                      padding: const EdgeInsets.only(
                                        top: 20,
                                        left: 20,
                                      ),
                                      child: ListView(
                                        children: options.map((c) {
                                          // user.currency holds canonical values like 'USD' or 'BIRR'
                                          final current =
                                              (user.currency ?? 'USD')
                                                  .toUpperCase();
                                          final isChecked =
                                              (current == 'USD' &&
                                                  c == getText('usd')) ||
                                              (current == 'BIRR' &&
                                                  c == getText('birr'));
                                          return ListTile(
                                            title: Text(c),
                                            trailing: isChecked
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
                                // Map localized label -> canonical currency code
                                final canonical = selected == getText('birr')
                                    ? 'BIRR'
                                    : 'USD';
                                context.read<UserAuthBloc>().add(
                                  UpdateCurrencyEvent(canonical),
                                );
                              }
                            },
                          ),
                          const Divider(height: 1),
                          SwitchRow(
                            activeColor: const Color.fromARGB(255, 204, 174, 27),
                            title: getText('notifications'),
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
                              context.read<UserAuthBloc>().add(
                                UpdateNotificationEvent(val),
                              );
                            },
                          ),
                        
                          const Divider(height: 1),
                          ThemeToggle(
                            activeColor: Colors.black,
                            title: getText('DarkMode'),
                            value: isDark,
                            onChanged: onDarkChanged,
                          ),
                          const Divider(height: 1),
                        ],
                      ),
                    ),
                    const SizedBox(height: 12),
                    // Version card
                    
                    const SizedBox(height: 12),
                    // Sign in / out action (UI only; uses existing bloc events)
                    Container(
                      decoration: cardDecoration(context),
                      child: ListTile(
                        leading: Icon(
                          user != null ? Icons.logout : Icons.login,
                          color: Colors.redAccent,
                        ),
                        title: Text(
                          user != null
                              ? getText('sign_out')
                              : getText('sign_in'),
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
            bottomNavigationBar: CurvedNavigationBar(
          index: 3, // Profile tab
          backgroundColor: Colors.transparent,
          color: Theme.of(context).cardColor,
          buttonBackgroundColor: Theme.of(context).cardColor,
          height: 60.0,
          items: [
            Icon(Icons.search, size: 30, color: Theme.of(context).textTheme.bodyLarge?.color),
            Icon(Icons.favorite_border, size: 30, color: Theme.of(context).textTheme.bodyLarge?.color),
            Icon(Icons.swap_horiz, size: 30, color: Theme.of(context).textTheme.bodyLarge?.color),
            Icon(Icons.person_outline, size: 30, color: Theme.of(context).textTheme.bodyLarge?.color),
          ],
          onTap: (index) {
            switch (index) {
              case 0:
                Navigator.pushReplacementNamed(context, '/chat');
                break;
              case 1:
                Navigator.pushReplacementNamed(context, '/saved');
                break;
              case 2:
                Navigator.pushReplacementNamed(context, '/compare-products');
                break;
              case 3:
              default:
                // Already on profile
                break;
            }
          },
            ),
          );
        },
      ),
    );
  }
}

String _initialFromId(String id) {
  final letters = id.trim();
  if (letters.isEmpty) return '';
  return letters.substring(0, 1).toUpperCase();
}
