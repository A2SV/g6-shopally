import 'package:flutter/material.dart';

Widget settingsRow({
  required String title,
  String? trailingText,
  VoidCallback? onTap,
}) {
  return ListTile(
    title: Text(title),
    trailing: Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        if (trailingText != null)
          Padding(
            padding: const EdgeInsets.only(right: 8.0),
            child: Text(trailingText),
          ),
        const Icon(Icons.chevron_right),
      ],
    ),
    onTap: onTap,
  );
}

Future<void> showSocialLoginBottomSheet(
  BuildContext context, {
  VoidCallback? onGoogle,
  VoidCallback? onApple,
}) {
  return showModalBottomSheet(
    context: context,
    backgroundColor: Theme.of(context).colorScheme.surface,
    isScrollControlled: true,

    shape: const RoundedRectangleBorder(
      borderRadius: BorderRadius.vertical(top: Radius.circular(16)),
    ),
    builder: (ctx) {
      return SizedBox(
        height: MediaQuery.of(ctx).size.height * 0.6,
        child: SafeArea(
          child: Padding(
            padding: const EdgeInsets.fromLTRB(16, 16, 16, 24),
            child: SizedBox(
              width: double.infinity,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: [
                  const SizedBox(height: 8),
                  Center(
                    child: Container(
                      width: 40,
                      height: 4,
                      decoration: BoxDecoration(
                        color: Theme.of(ctx).dividerColor.withOpacity(0.2),
                        borderRadius: BorderRadius.circular(2),
                      ),
                    ),
                  ),
                  const SizedBox(height: 40),
                  Text(
                    "Social Login",
                    textAlign: TextAlign.center,
                    style: Theme.of(ctx)
                        .textTheme
                        .headlineSmall
                        ?.copyWith(fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 6),
                  Text(
                    "Make a login using social network account",
                    textAlign: TextAlign.center,
                    style: Theme.of(ctx).textTheme.titleMedium,
                  ),
                  const SizedBox(height: 70),
                  ElevatedButton(
                    onPressed: () {
                      if (onGoogle != null) onGoogle();
                      Navigator.pop(ctx);
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Theme.of(ctx).colorScheme.surface,
                      foregroundColor: Theme.of(ctx).colorScheme.onSurface,
                      padding: const EdgeInsets.symmetric(
                        vertical: 18,
                        horizontal: 12,
                      ),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(10),
                        side: BorderSide(
                          color: Theme.of(ctx).dividerColor.withOpacity(0.07),
                        ),
                      ),
                    ),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Image.asset(
                          'assets/icon/googleicon.png',
                          width: 20,
                          height: 20,
                          fit: BoxFit.cover,
                        ),
                        const SizedBox(width: 10),
                        Text(
                          'Sign in with Google Account',
                          style: Theme.of(ctx)
                              .textTheme
                              .titleMedium
                              ?.copyWith(fontWeight: FontWeight.w700),
                        ),
                      ],
                    ),
                  ),
                  const SizedBox(height: 10),
                  ElevatedButton(
                    onPressed: () {
                      if (onApple != null) onApple();
                      Navigator.pop(ctx);
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Theme.of(ctx).colorScheme.surface,
                      foregroundColor: Theme.of(ctx).colorScheme.onSurface,
                      padding: const EdgeInsets.symmetric(
                        vertical: 18,
                        horizontal: 12,
                      ),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(10),
                        side: BorderSide(
                          color: Theme.of(ctx).dividerColor.withOpacity(0.07),
                        ),
                      ),
                    ),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Image.asset(
                          'assets/icon/apple.png',
                          width: 20,
                          height: 20,
                          fit: BoxFit.cover,
                        ),
                        const SizedBox(width: 10),
                        Text(
                          'Sign in with Apple Account',
                          style: Theme.of(ctx)
                              .textTheme
                              .titleMedium
                              ?.copyWith(fontWeight: FontWeight.w700),
                        ),
                      ],
                    ),
                  ),
                  const SizedBox(height: 40),
                ],
              ),
            ),
          ),
        ),
      );
    },
  );
}

Future<void> showSocialLoginDialog(
  BuildContext context, {
  VoidCallback? onGoogle,
  VoidCallback? onApple,
}) {
  return showDialog(
    context: context,
    builder: (ctx) =>
        socialLoginDialog(ctx, onGoogle: onGoogle, onApple: onApple),
  );
}

Widget SwitchRow({
  required String title,
  required bool value,
  required ValueChanged<bool> onChanged,
  required Color activeColor,
  
  bool dense = false,
}) {
  return ListTile(
    dense: dense,  
    title: Text(title),
    trailing: Switch.adaptive(
      value: value,  
      onChanged: onChanged,
      activeTrackColor: activeColor ,
    ),
    onTap: () => onChanged(!value),
  );
}



Widget ThemeToggle({
  required String title,
  required bool value,
  required ValueChanged<bool> onChanged,
  required Color activeColor,

  bool dense = false,
}) {
  return ListTile(
    dense: dense,
    title: Text(title),
    trailing: Switch.adaptive(
      value: value,
      onChanged: onChanged,
      activeTrackColor: activeColor,
    ),
    onTap: () => onChanged(!value),
  );
}

Widget QuickTile({
  required IconData icon,
  required String label,
  required VoidCallback onTap,
}) {
  return Builder(
    builder: (context) => Container(
      width: 150,
      decoration: BoxDecoration(
        color: Theme.of(context).cardColor,
        borderRadius: BorderRadius.circular(16),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.1),
            blurRadius: 4,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      child: InkWell(
        borderRadius: BorderRadius.circular(16),
        onTap: onTap,
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 14.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(icon),
              const SizedBox(width: 8),
              Text(label, style: const TextStyle(fontWeight: FontWeight.w500)),
            ],
          ),
        ),
      ),
    ),
  );
}

Widget PickerSheet({
  required BuildContext context,
  required String title,
  required List<String> options,
}) {
  return SafeArea(
    child: Column(
      mainAxisSize: MainAxisSize.min,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.fromLTRB(16, 8, 16, 8),
          child: Text(title, style: Theme.of(context).textTheme.titleMedium),
        ),
        const Divider(height: 1),
        ...options.map(
          (option) => ListTile(
            title: Text(option),
            onTap: () => Navigator.of(context).pop(option),
          ),
        ),
        const SizedBox(height: 8),
      ],
    ),
  );
}

BoxDecoration cardDecoration(BuildContext context) {
  return BoxDecoration(
    color: Theme.of(context).cardColor,
    borderRadius: BorderRadius.circular(16),
    boxShadow: const [
      BoxShadow(color: Color(0x11000000), blurRadius: 8, offset: Offset(0, 2)),
    ],
  );
}

Widget avatar({
  required String? name,
  required String? photoUrl,
  required String fallbackInitial,
}) {
  if (photoUrl != null && photoUrl.isNotEmpty) {
    return CircleAvatar(radius: 48, backgroundImage: NetworkImage(photoUrl));
  }

  final initial = (name != null && name.isNotEmpty)
      ? name.trim().substring(0, 1).toUpperCase()
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

Future<bool?> showSignOutDialog(BuildContext context) {
  return showDialog<bool>(
    context: context,
    builder: (ctx) {
      return AlertDialog(
        backgroundColor: Theme.of(ctx).colorScheme.surface,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
        titlePadding: const EdgeInsets.fromLTRB(24, 20, 24, 0),
        contentPadding: const EdgeInsets.fromLTRB(24, 12, 24, 0),
        actionsPadding: const EdgeInsets.fromLTRB(16, 8, 16, 12),
        title: Row(
          children: [
            Container(
              decoration: const BoxDecoration(
                color: Color(0x1FFF5252),
                shape: BoxShape.circle,
              ),
              padding: const EdgeInsets.all(8),
              child: const Icon(Icons.logout, color: Colors.redAccent),
            ),
            const SizedBox(width: 12),
            const Text('Sign out?'),
          ],
        ),
        content: const Text(
          'Are you sure you want to sign out of your account?',
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(ctx).pop(false),
            child: Text('Cancel', style: TextStyle(color: Theme.of(ctx).colorScheme.onSurface)),
          ),
          FilledButton(
            style: FilledButton.styleFrom(
              backgroundColor: Colors.redAccent,
              foregroundColor: Colors.white,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(10),
              ),
            ),
            onPressed: () => Navigator.of(ctx).pop(true),
            child: const Text('Sign out'),
          ),
        ],
      );
    },
  );
}

Widget socialLoginDialog(
  BuildContext context, {
  VoidCallback? onGoogle,
  VoidCallback? onApple,
}) {
  return Dialog(
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
    child: Padding(
      padding: const EdgeInsets.all(16.0),
      child: SizedBox(
        height: 500,
        width: double.infinity,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            const SizedBox(height: 90),
            Text(
              "Social Login",
              style: Theme.of(context)
                  .textTheme
                  .headlineSmall
                  ?.copyWith(fontSize: 30, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 5),
            Padding(
              padding: const EdgeInsets.only(left: 20),
              child: Text.rich(
                TextSpan(
                  text: "Make a login using social ",
                  style: Theme.of(context).textTheme.titleMedium,
                  children: [
                    WidgetSpan(
                      child: Padding(
                        padding: const EdgeInsets.only(left: 25),
                        child: Text(
                          "network account",
                          style: Theme.of(context).textTheme.titleMedium,
                        ),
                      ),
                    ),
                  ],
                ),
                textAlign: TextAlign.start,
              ),
            ),
            const SizedBox(height: 60),
            ElevatedButton(
              onPressed: () {
                if (onGoogle != null) onGoogle();
                Navigator.pop(context);
              },
              child: Row(
                children: [
                  Image.asset(
                    'assets/icon/googleicon.png',
                    width: 20,
                    height: 20,
                    fit: BoxFit.cover,
                  ),
                  const SizedBox(width: 10),
                  Text(
                    'Sign in with Google Account',
                    style: Theme.of(context)
                        .textTheme
                        .titleMedium
                        ?.copyWith(fontWeight: FontWeight.w700),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 10),
            ElevatedButton(
              onPressed: () {
                if (onApple != null) onApple();
                Navigator.pop(context);
              },
              child: Row(
                children: [
                  Image.asset(
                    'assets/icon/apple.png',
                    width: 20,
                    height: 20,
                    fit: BoxFit.cover,
                  ),
                  const SizedBox(width: 10),
                  Text(
                    'Sign in with Apple Account',
                    style: Theme.of(context)
                        .textTheme
                        .titleMedium
                        ?.copyWith(fontWeight: FontWeight.w700),
                   ),
                ],
              ),
            ),
          ],
        ),
      ),
    ),
  );
}
