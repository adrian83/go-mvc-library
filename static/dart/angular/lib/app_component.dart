



import 'package:angular2/core.dart';
import 'package:angular2/router.dart';

import 'author_service.dart';

import 'authors_list_comp.dart';
import 'b_comp.dart';

@Component(
    selector: 'my-app',
    template: '''
      <h1>{{title}}</h1>
      <a [routerLink]="['AuthorsListC']">acomp</a>
      <br/>
      <a [routerLink]="['BComp']">bcomp</a>
      <router-outlet></router-outlet>''',
    directives: const [ROUTER_DIRECTIVES], // components
    providers: const [AuthorService,ROUTER_PROVIDERS]) // services
@RouteConfig(const [
  const Route(path: '/authors', name: 'AuthorsListC', component: AuthorsListComponent),
  const Route(path: '/books', name: 'BComp', component: BComponent)
])
class AppComponent {
  String title = 'Tour of Heroes';
  AuthorService _authorService;

  AppComponent(this._authorService);
}
