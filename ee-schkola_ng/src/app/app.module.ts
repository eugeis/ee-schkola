import {NgModule} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
import {Http, HttpModule} from '@angular/http';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {TranslateLoader, TranslateModule} from '@ngx-translate/core';
import {TranslateHttpLoader} from '@ngx-translate/http-loader';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {AuthGuard} from './shared';
import {AlertService} from './shared/services/alert.service';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatOptionModule} from '@angular/material/core';
import {MatIconModule} from '@angular/material/icon';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatExpansionModule} from '@angular/material/expansion';
import {FamilyComponent} from './shared/ee/component/family/family.component';
import {EducationComponent} from './shared/ee/component/education/education.component';
import {LocationComponent} from './shared/ee/component/location/location.component';
import {PersonNameComponent} from './shared/ee/component/personname/personname.component';
import {ContactComponent} from './shared/ee/component/contact/contact.component';
import {ChurchInfoComponent} from './shared/ee/component/churchinfo/churchinfo.component';
import {AddressComponent} from './shared/ee/component/address/address.component';
import {TemplateComponent} from './shared/ee/template/components/example/template.component';
import {AlertComponent} from './layout/bs-component/components';

// AoT requires an exported function for factories
export function HttpLoaderFactory(http: HttpClient) {
    // for development
    // return new TranslateHttpLoader(http, '/start-angular/SB-Admin-BS4-Angular-4/master/dist/assets/i18n/', '.json');
    return new TranslateHttpLoader(http, '/assets/i18n/', '.json');
}

@NgModule({
    declarations: [
        AppComponent,
        FamilyComponent,
        EducationComponent,
        LocationComponent,
        PersonNameComponent,
        ContactComponent,
        ChurchInfoComponent,
        AddressComponent,
        TemplateComponent,
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        FormsModule,
        HttpModule,
        AppRoutingModule,
        TranslateModule.forRoot({
            loader: {
                provide: TranslateLoader,
                useFactory: HttpLoaderFactory,
                deps: [Http]
            }
        }),
        MatInputModule,
        MatFormFieldModule,
        MatSelectModule,
        MatOptionModule,
        MatIconModule,
        MatButtonModule,
        ReactiveFormsModule,
        MatExpansionModule,
    ],
    providers: [AuthGuard, AlertService],
    bootstrap: [AppComponent]
})
export class AppModule {
}
