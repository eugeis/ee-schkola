import { DatePipe } from '@angular/common';
import { Pipe, PipeTransform } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';

@Pipe({
    name: 'DateTimeTranslationPipe',
    pure: false
})
export class DateTimeTranslationPipePipe implements PipeTransform {
    constructor(private translateService: TranslateService) {}

    transform(value: Date, pattern: string = 'fullDate'): any {
        const datePipe: DatePipe = new DatePipe(
            this.translateService.currentLang ? this.translateService.currentLang : this.translateService.getDefaultLang());
        return datePipe.transform(value, pattern);
    }
}
