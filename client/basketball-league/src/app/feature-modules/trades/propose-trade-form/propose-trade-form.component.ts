import { trigger, transition, style, animate, state } from '@angular/animations';
import { AfterViewInit, Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialogRef, MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faMinus, faPlus } from '@fortawesome/free-solid-svg-icons';
import { AssetChoosingFormComponent } from '../asset-choosing-form/asset-choosing-form.component';
import { of, Observable, map, Subject, startWith, takeUntil, take } from 'rxjs';
import { MatSelect } from '@angular/material/select';

@Component({
  selector: 'app-propose-trade-form',
  templateUrl: './propose-trade-form.component.html',
  styleUrls: ['./propose-trade-form.component.css'],
  animations: [
      trigger("fadeIn", [
        transition(":enter", [
            style({ opacity: 0, transform: "translateX(-40px)" }),
            animate(
                "0.5s ease",
                style({ opacity: 1, transform: "translateX(0)" }),
            ),
        ]),
      ]),
      trigger('buttonState', [
        state('clicked', style({
          transform: 'scale(0.9)',
          opacity: 0.5
        })),
        transition('* => clicked', [
          animate('200ms')
        ]),
        transition('clicked => idle', [
          animate('200ms')
        ])
      ]),
  ],
})
export class ProposeTradeFormComponent implements OnInit, AfterViewInit, OnDestroy{
  buttonState: string = 'idle';
  removeAssetButtonState: string = 'idle';
  addPartnersAssetButtonState: string = 'idle';
  addYoursAssetButtonState: string = 'idle';
  focused: string = '';
  private ownDialogRef: any;
  public teams: string[] = [
    'Brooklyn Nets',
    'Golden State Warriors',
    'Los Angeles Lakers',
    'Los Angeles Clippers',
    'New Orleans Pelicans',
    'New York Knicks',
    'Oklahoma City Thunder',
    'San Antonio Spurs',
    'Boston Celtics',
    'Denver Nuggets',
    'Minnesota Timberwolves',
    'Cleveland Cavaliers',
    'Philadelphia 76ers',
    'Phoenix Suns',
    'Sacramento Kings',
    'Indiana Pacers',
    'Dallas Mavericks',
    'Miami Heat',
    'Orlando Magic',
    'Chicago Bulls',
    'Atlanta Hawks',
    'Toronto Raptors',
    'Charlotte Hornets',
    'Washington Wizards',
    'Detroit Pistons',
    'Utah Jazz',
    'Houston Rockets',
    'Memphis Grizzlies',
    'Portland Trail Blazers',
    'Milwaukee Bucks'
  ];

  public teamCtrl: FormControl<string | null> = new FormControl<string | null>('');
  public teamFilterCtrl: FormControl<string | null> = new FormControl<string | null>('');

  public filteredTeams: Observable<string[]> = of(this.teams);

  @ViewChild('singleSelect', { static: true }) singleSelect: MatSelect | undefined;

  protected _onDestroy = new Subject<void>();

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ProposeTradeFormComponent>,
              private dialogRefAsset: MatDialogRef<AssetChoosingFormComponent>,
              private dialog: MatDialog,) {
  }

  ngOnInit(): void {
    this.teamCtrl.setValue('');

    this.filteredTeams = this.teamFilterCtrl.valueChanges.pipe(
      startWith(''),
      map(value => this.filterTeams(value))
    );
  }

  private filterTeams(value: any): string[] {
    const filterValue = value?.toLowerCase() || '';
    return this.teams.filter(team => team.toLowerCase().includes(filterValue));
  }

  ngAfterViewInit() {
    this.setInitialValue();
  }

  ngOnDestroy() {
    this._onDestroy.next();
    this._onDestroy.complete();
  }

  protected setInitialValue() {
    this.filteredTeams
      .pipe(take(1), takeUntil(this._onDestroy))
      .subscribe(() => {
        this.singleSelect!.compareWith = (a: string, b: string) => a.toLowerCase() === b.toLowerCase();
      });
  }

  // Izmeni samo da bude trejd kao
  proposeTradeForm = new FormGroup({
    name: new FormControl('', [Validators.required]),
    description: new FormControl('', [Validators.required]),
    //duration: new FormControl('', [Validators.required]),
    occurrenceTime: new FormControl(null, [Validators.required]),
    occurrenceDate: new FormControl(null, [Validators.required]),
    //guide: new FormControl('', [Validators.required]),
    capacity: new FormControl('', [Validators.required]),
    picturePath: new FormControl('', [Validators.required]),
    category: new FormControl('', [Validators.required]),
  });

  addTourButtonClicked() {
    // const selectedCategoryString: string = this.addTourForm.value.category ?? '';
    // let selectedCategory: TourCategory;

    // switch (selectedCategoryString) {
    //   case 'ART_COLLECTIONS':
    //     selectedCategory = TourCategory.ArtCollections;
    //     break;
    //   case 'HISTORICAL_EXHIBITS':
    //     selectedCategory = TourCategory.HistoricalExhibits;
    //     break;
    //   case 'SCIENCE_AND_TECHNOLOGY':
    //     selectedCategory = TourCategory.ScienceAndTechnology;
    //     break;
    //   case 'CULTURAL_HERITAGE':
    //     selectedCategory = TourCategory.CulturalHeritage;
    //     break;
    //   case 'ANCIENT_ART':
    //     selectedCategory = TourCategory.AncientArt;
    //     break;
    //   case 'EUROPEAN_PAINTINGS':
    //     selectedCategory = TourCategory.EuropeanPaintings;
    //     break;
    //   case 'MODERN_ART':
    //     selectedCategory = TourCategory.ModernArt;
    //     break;
    //   case 'AMERICAN_ART':
    //     selectedCategory = TourCategory.AmericanArt;
    //     break;
    //   case 'ASIAN_ART':
    //     selectedCategory = TourCategory.AsianArt;
    //     break;
    //   case 'AFRICAN_CULTURE':
    //     selectedCategory = TourCategory.AfricanCulture;
    //     break;
    //   case 'ISLAMIC_ART':
    //     selectedCategory = TourCategory.IslamicArt;
    //     break;
    //   case 'COSTUME_INSTITUTE':
    //     selectedCategory = TourCategory.CostumeInstitute;
    //     break;
    //   case 'ARMS_AND_ARMOR':
    //     selectedCategory = TourCategory.ArmsAndArmor;
    //     break;
    //   default:
    //     console.error("Invalid category selected.");
    //     return;
    // }

    // const tour: Tour = {
    //   name: this.addTourForm.value.name || "",
    //   description: this.addTourForm.value.description || "",
    //   occurrenceDateTime: this.addTourForm.value.occurrenceDate || new Date(),
    //   adultTicketPrice: this.adultTicketPrice || "",
    //   minorTicketPrice: this.minorTicketPrice || "",
    //   capacity: this.addTourForm.value.capacity || "",
    //   picturePath: this.addTourForm.value.picturePath || "",
    //   category: selectedCategory,
    // };

    // console.log(tour);

    // if (this.addTourForm.valid) {
    //     this.buttonState = 'clicked';
    //     setTimeout(() => { this.buttonState = 'idle'; }, 200);

    //     // Postavi datum i vreme
    //     const dateValue: Date | null = this.addTourForm.value.occurrenceDate!;
    //     const timeValue: string | null = this.addTourForm.value.occurrenceTime!;

    //     const [hours, minutes] = (timeValue as string).split(':');
    //     const dateTime = new Date(dateValue);
    //     dateTime.setHours(Number(hours) + 1);
    //     dateTime.setMinutes(Number(minutes));

    //     const d = new Date(dateValue);
    //     d.setHours(Number(hours));
    //     d.setMinutes(Number(minutes));

    //     tour.occurrenceDateTime = dateTime;

    //     if(this.selectedCurator.length != 0){
    //       tour.guideId = this.selectedCurator[0].id;
    //       if(this.selectedExhibitions.length != 0){
    //         tour.duration = (this.selectedExhibitions.length * 15).toString();
    //         tour.exhibitions = this.selectedExhibitions;
    //         this.toursService.addTour(tour).subscribe({
    //           next: () => {
    //             this.showNotification('Tour successfully added!')
    //             this.dialogRef.close();
    //           },
    //         });
    //       }
    //       else{
    //         this.showNotification('Please select at least one exhibition')
    //       }
    //     }
    //     else{
    //       this.showNotification('Please select a curator')
    //     }
    // }
    // else{
    //   this.showNotification('Please fill out the form correctly')
    // }
  }

  addPartnersAssetButtonClicked(): void {
    this.addPartnersAssetButtonState = 'clicked';
    setTimeout(() => { this.addPartnersAssetButtonState = 'idle'; }, 200);
    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        // TODO: Ovde treba proslediti tim da se zna cija imovina da se prikaze
    });

    if (this.dialogRefAsset) {
      this.dialogRefAsset.afterClosed().subscribe((result: any) => {
        // TODO: Ovde treba dodati osvezavanje liste kada se odabere nova imovina za trejdovanje
      });
    }
  }

  addYoursAssetButtonClicked(): void {
    this.addYoursAssetButtonState = 'clicked';
    setTimeout(() => { this.addYoursAssetButtonState = 'idle'; }, 200);
    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        // TODO: Ovde treba proslediti tim da se zna cija imovina da se prikaze
    });

    if (this.dialogRefAsset) {
      this.dialogRefAsset.afterClosed().subscribe((result: any) => {
        // TODO: Ovde treba dodati osvezavanje liste kada se odabere nova imovina za trejdovanje
      });
    }
  }

  removeAssetButtonClicked(): void {
    // Ovde samo proveri da li je sve okej sto se tice same animacije
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  overviewClicked(){
    this.dialogRef.close();
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
