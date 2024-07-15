import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialogRef, MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { ProposeTradeFormComponent } from '../propose-trade-form/propose-trade-form.component';
import { trigger, transition, style, animate, state } from '@angular/animations';

@Component({
  selector: 'app-asset-choosing-form',
  templateUrl: './asset-choosing-form.component.html',
  styleUrls: ['./asset-choosing-form.component.css'],
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
export class AssetChoosingFormComponent implements OnInit{
  finishButtonState: string = 'idle';
  removeAssetButtonState: string = 'idle';
  focused: string = '';
  private ownDialogRef: any;

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ProposeTradeFormComponent>,
              private dialogRefAsset: MatDialogRef<AssetChoosingFormComponent>,
              private dialog: MatDialog,) {
  }

  ngOnInit(): void {
    
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

  removeAssetButtonClicked(): void {
    // Ovde samo proveri da li je sve okej sto se tice same animacije
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);
  }

  finishButtonClicked(): void {
    this.finishButtonState = 'clicked';
    setTimeout(() => { this.finishButtonState = 'idle'; }, 200);
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
