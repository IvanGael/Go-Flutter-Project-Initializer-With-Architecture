package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

const (
	BLoC              string = "BLoC (Business Logic Component)"
	Provider          string = "Provider"
	Redux             string = "Redux"
	ScopedModel       string = "Scoped Model"
	Mvvm              string = "MVVM (Model-View-ViewModel)"
	Mvc               string = "MVC (Model-View-Controller)"
	Cubit             string = "Cubit"
	Riverpod          string = "Riverpod"
	GetX              string = "GetX"
	MobX              string = "MobX"
	StatesRebuilder   string = "States Rebuilder"
	CleanArchitecture string = "Clean Architecture"
)

func main() {
	var architecture string

	// List of architectures
	architectures := []string{
		BLoC,
		Provider,
		Redux,
		ScopedModel,
		Mvvm,
		Mvc,
		Cubit,
		Riverpod,
		GetX,
		MobX,
		StatesRebuilder,
		CleanArchitecture,
	}

	// Prompt the user to select an architecture
	prompt := &survey.Select{
		Message: "Choose the architecture you want to use for your Flutter project:",
		Options: architectures,
	}
	survey.AskOne(prompt, &architecture)

	var projectName string
	// Prompt the user to enter the project name
	promptInput := &survey.Input{
		Message: "Enter the project name:",
	}
	survey.AskOne(promptInput, &projectName)

	var path string
	// Prompt the user to enter the path
	promptInput = &survey.Input{
		Message: "Enter the path to initialize the project (press Enter for current directory):",
	}
	survey.AskOne(promptInput, &path)

	path = strings.TrimSpace(path)
	if path == "" {
		path = "."
	}

	// Create the Flutter project
	initializeProject(architecture, projectName, path)
}

func initializeProject(architecture, projectName, path string) {
	fmt.Printf("Initializing project %s using %s architecture in %s...\n", projectName, architecture, path)

	// Create the Flutter project
	cmd := exec.Command("flutter", "create", projectName)
	cmd.Dir = path
	executeCommand(cmd)

	projectPath := filepath.Join(path, projectName)

	// Add architecture-specific packages and example classes
	switch architecture {
	case BLoC:
		addBlocArchitecture(projectPath)
	case Provider:
		addProviderArchitecture(projectPath)
	case Redux:
		addReduxArchitecture(projectPath)
	case ScopedModel:
		addScopedModelArchitecture(projectPath)
	case Mvvm:
		addMvvmArchitecture(projectPath)
	case Mvc:
		addMvcArchitecture(projectPath)
	case Cubit:
		addCubitArchitecture(projectPath)
	case Riverpod:
		addRiverpodArchitecture(projectPath)
	case GetX:
		addGetXArchitecture(projectPath)
	case MobX:
		addMobXArchitecture(projectPath)
	case StatesRebuilder:
		addStatesRebuilderArchitecture(projectPath)
	case CleanArchitecture:
		addCleanArchitecture(projectPath)
	default:
		fmt.Printf("Architecture %s is not supported yet.\n", architecture)
		return
	}

	fmt.Printf("Project %s initialized successfully with %s architecture in %s\n", projectName, architecture, path)
}

func addBlocArchitecture(projectPath string) {
	//Add necessary packages for BLoC
	cmd := exec.Command("flutter", "pub", "add", "flutter_bloc", "bloc")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'bloc/counter_bloc.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: BlocProvider(
        create: (context) => CounterBloc(),
        child: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterBloc = BlocProvider.of<CounterBloc>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter BLoC Example"),
      ),
      body: Center(
        child: BlocBuilder<CounterBloc, int>(
          builder: (context, count) {
            return Text("$count");
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          counterBloc.add(CounterEvent.increment);
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterBlocContent := `
import 'package:bloc/bloc.dart';

enum CounterEvent { increment }

class CounterBloc extends Bloc<CounterEvent, int> {
  CounterBloc() : super(0);

  @override
  Stream<int> mapEventToState(CounterEvent event) async* {
    switch (event) {
      case CounterEvent.increment:
        yield state + 1;
        break;
    }
  }
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "bloc"), 0755)
	createFile(filepath.Join(projectPath, "lib", "bloc", "counter_bloc.dart"), counterBlocContent)
}

func addProviderArchitecture(projectPath string) {
	// Add necessary packages for Provider
	cmd := exec.Command("flutter", "pub", "add", "provider")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'provider/counter_provider.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => CounterProvider()),
      ],
      child: MaterialApp(
        home: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterProvider = Provider.of<CounterProvider>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter Provider Example"),
      ),
      body: Center(
        child: Text("${counterProvider.count}"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          counterProvider.increment();
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterProviderContent := `
import 'package:flutter/material.dart';

class CounterProvider with ChangeNotifier {
  int _count = 0;

  int get count => _count;

  void increment() {
    _count++;
    notifyListeners();
  }
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "provider"), 0755)
	createFile(filepath.Join(projectPath, "lib", "provider", "counter_provider.dart"), counterProviderContent)
}

func addReduxArchitecture(projectPath string) {
	// Add necessary packages for Redux
	cmd := exec.Command("flutter", "pub", "add", "redux", "flutter_redux")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';
import 'redux/counter_reducer.dart';

void main() {
  final store = Store<int>(counterReducer, initialState: 0);
  runApp(MyApp(store: store));
}

class MyApp extends StatelessWidget {
  final Store<int> store;

  MyApp({required this.store});

  @override
  Widget build(BuildContext context) {
    return StoreProvider<int>(
      store: store,
      child: MaterialApp(
        home: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter Redux Example"),
      ),
      body: Center(
        child: StoreConnector<int, String>(
          converter: (store) => store.state.toString(),
          builder: (context, count) {
            return Text(count);
          },
        ),
      ),
      floatingActionButton: StoreConnector<int, VoidCallback>(
        converter: (store) {
          return () => store.dispatch(CounterAction.increment);
        },
        builder: (context, callback) {
          return FloatingActionButton(
            onPressed: callback,
            child: Icon(Icons.add),
          );
        },
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterReducerContent := `
enum CounterAction { increment }

int counterReducer(int state, dynamic action) {
  if (action == CounterAction.increment) {
    return state + 1;
  }
  return state;
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "redux"), 0755)
	createFile(filepath.Join(projectPath, "lib", "redux", "counter_reducer.dart"), counterReducerContent)
}

func addScopedModelArchitecture(projectPath string) {
	// Add necessary packages for ScopedModel
	cmd := exec.Command("flutter", "pub", "add", "scoped_model")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:scoped_model/scoped_model.dart';
import 'scoped_model/counter_model.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ScopedModel<CounterModel>(
      model: CounterModel(),
      child: MaterialApp(
        home: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter ScopedModel Example"),
      ),
      body: Center(
        child: ScopedModelDescendant<CounterModel>(
          builder: (context, child, model) {
            return Text("${model.count}");
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          ScopedModel.of<CounterModel>(context).increment();
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterModelContent := `
import 'package:scoped_model/scoped_model.dart';

class CounterModel extends Model {
  int _count = 0;

  int get count => _count;

  void increment() {
    _count++;
    notifyListeners();
  }
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "scoped_model"), 0755)
	createFile(filepath.Join(projectPath, "lib", "scoped_model", "counter_model.dart"), counterModelContent)
}

func addMvvmArchitecture(projectPath string) {
	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'viewmodel/counter_viewmodel.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => CounterViewModel()),
      ],
      child: MaterialApp(
        home: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterViewModel = Provider.of<CounterViewModel>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter MVVM Example"),
      ),
      body: Center(
        child: Text("${counterViewModel.count}"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          counterViewModel.increment();
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterViewModelContent := `
import 'package:flutter/material.dart';

class CounterViewModel extends ChangeNotifier {
  int _count = 0;

  int get count => _count;

  void increment() {
    _count++;
    notifyListeners();
  }
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "viewmodel"), 0755)
	createFile(filepath.Join(projectPath, "lib", "viewmodel", "counter_viewmodel.dart"), counterViewModelContent)
}

func addMvcArchitecture(projectPath string) {
	// Add necessary packages for MVC
	cmd := exec.Command("flutter", "pub", "add", "mvc_pattern")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:mvc_pattern/mvc_pattern.dart';
import 'controller/counter_controller.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  State createState() => _MyHomePageState();
}

class _MyHomePageState extends StateMVC<MyHomePage> {
  _MyHomePageState() : super(CounterController()) {
    con = controller as CounterController;
  }

  late CounterController con;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter MVC Example"),
      ),
      body: Center(
        child: Text("${con.count}"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: con.increment,
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterControllerContent := `
import 'package:mvc_pattern/mvc_pattern.dart';

class CounterController extends ControllerMVC {
  int _count = 0;

  int get count => _count;

  void increment() {
    setState(() {
      _count++;
    });
  }
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "controller"), 0755)
	createFile(filepath.Join(projectPath, "lib", "controller", "counter_controller.dart"), counterControllerContent)
}

func addCubitArchitecture(projectPath string) {
	// Add necessary packages for Cubit
	cmd := exec.Command("flutter", "pub", "add", "flutter_bloc", "bloc")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'cubit/counter_cubit.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: BlocProvider(
        create: (context) => CounterCubit(),
        child: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterCubit = BlocProvider.of<CounterCubit>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter Cubit Example"),
      ),
      body: Center(
        child: BlocBuilder<CounterCubit, int>(
          builder: (context, count) {
            return Text("$count");
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          counterCubit.increment();
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterCubitContent := `
import 'package:bloc/bloc.dart';

class CounterCubit extends Cubit<int> {
  CounterCubit() : super(0);

  void increment() => emit(state + 1);
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "cubit"), 0755)
	createFile(filepath.Join(projectPath, "lib", "cubit", "counter_cubit.dart"), counterCubitContent)
}

func addRiverpodArchitecture(projectPath string) {
	// Add necessary packages for Riverpod
	cmd := exec.Command("flutter", "pub", "add", "flutter_riverpod")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

void main() {
  runApp(ProviderScope(child: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: MyHomePage(),
    );
  }
}

final counterProvider = StateProvider<int>((ref) {
  return 0;
});

class MyHomePage extends ConsumerWidget {
  @override
  Widget build(BuildContext context, ScopedReader watch) {
    final count = watch(counterProvider).state;
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter Riverpod Example"),
      ),
      body: Center(
        child: Text("$count"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          context.read(counterProvider).state++;
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)
}

func addGetXArchitecture(projectPath string) {
	// Add necessary packages for GetX
	cmd := exec.Command("flutter", "pub", "add", "get")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'controller/counter_controller.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final CounterController counterController = Get.put(CounterController());
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter GetX Example"),
      ),
      body: Center(
        child: Obx(() {
          return Text("${counterController.count}");
        }),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: counterController.increment,
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterControllerContent := `
import 'package:get/get.dart';

class CounterController extends GetxController {
  var count = 0.obs;

  void increment() => count++;
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "controller"), 0755)
	createFile(filepath.Join(projectPath, "lib", "controller", "counter_controller.dart"), counterControllerContent)
}

func addMobXArchitecture(projectPath string) {
	// Add necessary packages for MobX
	cmd1 := exec.Command("flutter", "pub", "add", "flutter_mobx", "mobx", "provider")
	cmd1.Dir = projectPath
	executeCommand(cmd1)

	cmd2 := exec.Command("flutter", "packages", "pub", "run", "build_runner", "build")
	cmd2.Dir = projectPath
	executeCommand(cmd2)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:flutter_mobx/flutter_mobx.dart';
import 'package:provider/provider.dart';
import 'store/counter_store.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        Provider<CounterStore>(create: (_) => CounterStore()),
      ],
      child: MaterialApp(
        home: MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterStore = Provider.of<CounterStore>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter MobX Example"),
      ),
      body: Center(
        child: Observer(
          builder: (_) => Text("${counterStore.count}"),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: counterStore.increment,
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	counterStoreContent := `
import 'package:mobx/mobx.dart';

part 'counter_store.g.dart';

class CounterStore = _CounterStore with _$CounterStore;

abstract class _CounterStore with Store {
  @observable
  int count = 0;

  @action
  void increment() {
    count++;
  }
}
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "store"), 0755)
	createFile(filepath.Join(projectPath, "lib", "store", "counter_store.dart"), counterStoreContent)
	createFile(filepath.Join(projectPath, "lib", "store", "counter_store.g.dart"), "")
}

func addStatesRebuilderArchitecture(projectPath string) {
	// Add necessary packages for States Rebuilder
	cmd := exec.Command("flutter", "pub", "add", "states_rebuilder")
	cmd.Dir = projectPath
	executeCommand(cmd)

	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:states_rebuilder/states_rebuilder.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatelessWidget {
  final counterRM = RM.inject(() => 0);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter States Rebuilder Example"),
      ),
      body: Center(
        child: OnBuilder(
          listenTo: counterRM,
          builder: () => Text("${counterRM.state}"),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          counterRM.state++;
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)
}

func addCleanArchitecture(projectPath string) {
	// Create example classes
	mainContent := `
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'injection_container.dart' as di;
import 'features/counter/presentation/pages/counter_page.dart';

void main() {
  di.init();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: di.providers,
      child: MaterialApp(
        home: CounterPage(),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "main.dart"), mainContent)

	injectionContainerContent := `
import 'package:get_it/get_it.dart';
import 'package:provider/provider.dart';
import 'features/counter/presentation/provider/counter_provider.dart';

final sl = GetIt.instance;

void init() {
  sl.registerFactory(() => CounterProvider());

  providers = [
    ChangeNotifierProvider(create: (_) => sl<CounterProvider>()),
  ];
}

List<SingleChildWidget> providers = [];
`
	os.MkdirAll(filepath.Join(projectPath, "lib", "features", "counter", "presentation", "pages"), 0755)
	os.MkdirAll(filepath.Join(projectPath, "lib", "features", "counter", "presentation", "provider"), 0755)
	createFile(filepath.Join(projectPath, "lib", "injection_container.dart"), injectionContainerContent)

	counterPageContent := `
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../provider/counter_provider.dart';

class CounterPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final counterProvider = Provider.of<CounterProvider>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text("Flutter Clean Architecture Example"),
      ),
      body: Center(
        child: Text("${counterProvider.count}"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: counterProvider.increment,
        child: Icon(Icons.add),
      ),
    );
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "features", "counter", "presentation", "pages", "counter_page.dart"), counterPageContent)

	counterProviderContent := `
import 'package:flutter/material.dart';

class CounterProvider with ChangeNotifier {
  int _count = 0;

  int get count => _count;

  void increment() {
    _count++;
    notifyListeners();
  }
}
`
	createFile(filepath.Join(projectPath, "lib", "features", "counter", "presentation", "provider", "counter_provider.dart"), counterProviderContent)
}

func executeCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func createFile(filePath, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
