import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TrainingProposalCardComponent } from './training-proposal-card.component';

describe('TrainingProposalCardComponent', () => {
  let component: TrainingProposalCardComponent;
  let fixture: ComponentFixture<TrainingProposalCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TrainingProposalCardComponent]
    });
    fixture = TestBed.createComponent(TrainingProposalCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
