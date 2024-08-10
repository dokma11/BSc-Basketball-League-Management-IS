import requests
import bs4
"""
    This script is mainly used to help me insert valid and realistic asset data into my database, the scraping doesn't 
    cover fully every single detail that needs to be inserted, but it still saves me tons of time
"""
team_abbreviations = {
            'Brooklyn Nets': 'BRK',
            'Golden State Warriors': 'GSW',
            'Los Angeles Lakers': 'LAL',
            'Los Angeles Clippers': 'LAC',
            'New Orleans Pelicans': 'NOP',
            'New York Knicks': 'NYK',
            'Oklahoma City Thunder': 'OKC',
            'San Antonio Spurs': 'SAS',
            'Boston Celtics': 'BOS',
            'Denver Nuggets': 'DEN',
            'Minnesota Timberwolves': 'MIN',
            'Cleveland Cavaliers': 'CLE',
            'Philadelphia 76ers': 'PHI',
            'Phoenix Suns': 'PHO',
            'Sacramento Kings': 'SAC',
            'Indiana Pacers': 'IND',
            'Dallas Mavericks': 'DAL',
            'Miami Heat': 'MIA',
            'Orlando Magic': 'ORL',
            'Chicago Bulls': 'CHI',
            'Atlanta Hawks': 'ATL',
            'Toronto Raptors': 'TOR',
            'Charlotte Hornets': 'CHA',
            'Washington Wizards': 'WAS',
            'Detroit Pistons': 'DET',
            'Utah Jazz': 'UTA',
            'Houston Rockets': 'HOU',
            'Memphis Grizzlies': 'MEM',
            'Portland Trail Blazers': 'POR',
            'Milwaukee Bucks': 'MIL',
}

base_url = 'https://www.basketball-reference.com/'


def scrape_players():
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation + '/2025.html#all_roster')
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        table = soup.find('table', {'id': 'roster'})

        trs = table.find_all('tr')

        insert_seq = ''
        for tr in trs:
            tds = tr.find_all('td')
            if len(tds) > 0:
                # Split into first name and last name
                name_parts = tds[0].find('a').get_text().split(' ')
                first_name = name_parts[0]
                if len(name_parts) > 2:
                    last_name = ' '.join(name_parts[1:])
                else:
                    last_name = name_parts[1]

                # There are no precise position information, so I have to scrape positions manually from player's page
                a_tag = tds[0].find('a')
                player_page = base_url + a_tag.get('href')
                player_page_response = requests.get(player_page)
                player_page_response.raise_for_status()
                player_page_soup = bs4.BeautifulSoup(player_page_response.text, 'html.parser')

                position_div = player_page_soup.find('div', {'id': 'meta'})
                p_tags = position_div.find_all('p')
                for p in p_tags:
                    if p.get_text().__contains__('Position'):
                        if p.get_text().__contains__('Point Guard') or p.get_text().__contains__('Guard'):
                            position = 'PG'

                        elif p.get_text().__contains__('Shooting Guard'):
                            position = 'SG'

                        elif p.get_text().__contains__('Small Forward') or p.get_text().__contains__('Forward'):
                            position = 'SF'

                        elif p.get_text().__contains__('Power Forward'):
                            position = 'PF'

                        elif p.get_text().__contains__('Center'):
                            position = 'C'

                # Convert height from feet and inches to centimeters
                height_parts = tds[2].get_text().split('-')     # Split the string for easier conversion
                height = int(int(height_parts[0]) * 30.48 + int(height_parts[1]) * 2.54)

                # Convert weight from pounds to kilograms
                if len(tds[3].get_text()) > 0:
                    weight = str(int(float(tds[3].get_text()) * 0.453592))
                else:
                    weight = 'Unknown'  # There is no information about the weight on the basketball-reference page

                # Convert birthday to the correct format
                birthday_parts = tds[4].get_text().split(',')
                date_parts = birthday_parts[0].split(' ')
                full_birthday = date_parts[1] + '-' + date_parts[0][:3].upper() + '-' + birthday_parts[1][1:]

                password = 'igrac123'
                email = first_name.lower() + '@gmail.com'
                user_insert = f"INSERT INTO KORISNIK VALUES (0, '{email}', '{first_name}', '{last_name}', '{full_birthday}', '{password}', 'Zaposleni')"
                contract_insert = f"INSERT INTO UGOVOR VALUES ({contract_id}, SYSDATE, SYSDATE, '100', 'NO_OPTION', 16, 0 )" # Last one is Player type contract
                employee_insert = f"INSERT INTO ZAPOSLENI VALUES ({contract_id}, 'Igrac', '100', {contract_id})"
                player_insert = f"INSERT INTO IGRAC VALUES ({contract_id}, '{height}', '{weight}', '{position}')"
                contract_id = contract_id + 1

                insert_seq = insert_seq + user_insert + ';\n' + contract_insert + ';\n' + employee_insert + ';\n' + player_insert + ';\n'

        with open('./res.txt', 'w', encoding='utf-8') as file:
            file.write(insert_seq)


def scrape_teams():
    insert_seq = ''
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation)
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        div_parent_tag = soup.find_all('div', {'id': 'meta'})
        for div in div_parent_tag:
            span_tags = div.find_all('span')
            for span in span_tags:
                if len(span.get_text()) > 0:
                    team_name = span.get_text()
                    print(team_name)
            p_tags = div.find_all('p')
            for p in p_tags:
                if p.get_text().__contains__('Location'):
                    team_location = p.get_text()
                if p.get_text().__contains__('Seasons'):
                    # Get the establishment year from the string
                    p_parts = p.get_text().split(';')
                    p_parts[1] = p_parts[1].strip().split()
                    team_establishment_year = p_parts[0].split('-')[0]

        team_insert = f"INSERT INTO TIM VALUES (0, '{team_name}', '{team_establishment_year}', '{team_location}')"
        insert_seq = insert_seq + team_insert + ';\n'
    with open('./res.txt', 'w', encoding='utf-8') as file:
        file.write(insert_seq)


def scrape_managers():
    insert_seq = ''
    contract_id = 0
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation + '/executives.html')
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        table = soup.find('table', _class='sortable stats_table now_sortable')
        a_tags = table.find_all('a')
        if a_tags:
            # Get the last <a> (the current manager is always the last one)
            manager_name_parts = a_tags[-1].get_text().split(' ')
            manager_first_name = manager_name_parts[0]
            manager_last_name = manager_name_parts[1]
            # Id's will be handled on server side
            password = 'menadzer123'
            email = manager_first_name.lower() + '@gmail.com'
            user_insert = f"INSERT INTO KORISNIK VALUES (0, '{email}', '{manager_first_name}', '{manager_last_name}', '{password}', 'Zaposleni')"
            contract_insert = f"INSERT INTO UGOVOR VALUES ({contract_id}, SYSDATE, SYSDATE, '100', 'NO_OPTION', 16, 1 )" # Last one is Manager type contract
            employee_insert = f"INSERT INTO ZAPOSLENI VALUES ({contract_id}, 'Menadzer', '200', {contract_id})"
            manager_insert = f"INSERT INTO MENADZER VALUES ({contract_id})"
            contract_id = contract_id + 1
            insert_seq = insert_seq + user_insert + ';\n' + contract_insert + ';\n' + employee_insert + ';\n' + manager_insert + ';\n'

    with open('./res.txt', 'w', encoding='utf-8') as file:
        file.write(insert_seq)


def scrape_coaches():
    insert_seq = ''
    contract_id = 0
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation)
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        table = soup.find('table', _class='sortable stats_table now_sortable sticky_table eql rel lel')
        tr_tags = table.find_all('tr')
        a_tags = tr_tags[0].find_all('a')
        if a_tags:
            # Manually go to the coaches page for more information
            coach_page = base_url + 'coaches/' + a_tags[-1].get('href')
            coach_page_response = requests.get(coach_page)
            coach_page_response.raise_for_status()
            coach_page_soup = bs4.BeautifulSoup(coach_page_response.text, 'html.parser')

            span_tags = coach_page_soup.find_all('span')
            coach_name_parts = span_tags[-1].get_text().split(' ')
            coach_first_name = coach_name_parts[0]
            coach_last_name = coach_name_parts[1]

            birthday_span_tag = coach_page_soup.find('span', {'id': 'necro-birth'})
            coach_birthday = birthday_span_tag.get_text()

            table = coach_page_soup.find('table', _class='suppress_glossary sortable stats_table now_sortable')
            th_tags = table.find_all('th')
            coach_years_of_experience = 0
            for th in th_tags:
                if len(th.get_text()) > 0:
                    th_parts = th.get_text().split(' ')
                    coach_years_of_experience += int(th_parts[0])

            password = 'trener123'
            email = coach_first_name.lower() + '@gmail.com'
            user_insert = f"INSERT INTO KORISNIK VALUES (0, '{email}', '{coach_first_name}', '{coach_last_name}', '{coach_birthday}' '{password}', 'Zaposleni')"
            contract_insert = f"INSERT INTO UGOVOR VALUES ({contract_id}, SYSDATE, SYSDATE, '100', 'NO_OPTION', 16, 2)" # Last one is Coach type contract
            employee_insert = f"INSERT INTO ZAPOSLENI VALUES ({contract_id}, 'Coach', '150', {contract_id})"
            coach_insert = f"INSERT INTO TRENER VALUES ({contract_id}, {coach_years_of_experience}, '')"
            contract_id = contract_id + 1
            insert_seq = insert_seq + user_insert + ';\n' + contract_insert + ';\n' + employee_insert + ';\n' + coach_insert + ';\n'

    with open('./res.txt', 'w', encoding='utf-8') as file:
        file.write(insert_seq)


if __name__ == "__main__":
    print('Scraping players... \n')
    scrape_players()
    print('Player scraping finished!\n')
    print('Scraping teams... \n')
    scrape_teams()
    print('Team scraping finished\n')
    print('Scraping managers... \n')
    scrape_managers()
    print('Manager scraping finished\n')
    print('Scraping coaches... \n')
    scrape_coaches()
    print('Coach scraping finished\n')
